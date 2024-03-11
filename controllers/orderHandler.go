package controllers

import (
	"asignrest/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrdersController struct {
	DB *gorm.DB
}
func NewOrderController(db *gorm.DB) *OrdersController {
	return &OrdersController{DB: db}
}

func (rc *OrdersController)CreateOrders(c *gin.Context)  {
	var (
		ordersReq models.OrdersRequest
	  order models.Orders
	)
	

	if err := c.ShouldBindJSON(&ordersReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order = models.Orders{
		Ordered_at: ordersReq.OrderedAt,
		Customer_name: ordersReq.CustomName,
	}

	for _, v := range ordersReq.Items {
		item := models.Items{
			Item_code: v.ItemCode,
			Description: v.Description,
			Quantity: v.Quantity,
		}
		// Menyimpan item ke basis data
		result := rc.DB.Create(&item)
		if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
				return
		}
	}
	
	result := rc.DB.Create(&order)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Berhasil tambah data",
		"orders":  order,
	})
}


func (rc *OrdersController)UpdateOrders(c *gin.Context)  {
	id := c.Param("id")
	var (
		ordersReq models.OrdersRequest
	  order models.Orders
	)

	if err := c.ShouldBindJSON(&ordersReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	checkOrder := rc.DB.First(&order, id)
	if checkOrder.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "database not found"})
		return
	}

	convertedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
			fmt.Println("Error:", err)
			return
	}

	order = models.Orders{
		Order_id: uint(convertedId),
		Ordered_at: ordersReq.OrderedAt,
		Customer_name: ordersReq.CustomName,
	}

	saveOrder := rc.DB.Save(&order)
	if saveOrder.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": saveOrder.Error.Error()})
		return
	}

	for _, v := range ordersReq.Items {
		item := models.Items{
			Item_id: v.LineItem_id,
			Item_code: v.ItemCode,
			Description: v.Description,
			Quantity: v.Quantity,
		}
		checkItems := rc.DB.First(&item, item.Item_id)
		if checkItems.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "database not found"})
			return 
		}
		// Menyimpan item ke basis data
		result := rc.DB.Save(&item)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Berhasil update data",
		"orders":  order,
	})
}



func (rc *OrdersController) GetAllOrders(c *gin.Context) {

	var orders []models.Orders
	result := rc.DB.Find(&orders)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (rc *OrdersController) GetOrdersByID(c *gin.Context) {
	id := c.Param("id")

	var order models.Orders
	result := rc.DB.First(&order, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (rc *OrdersController) DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	fmt.Println("mencoba")
	var order models.Orders
	result := rc.DB.First(&order, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	result = rc.DB.Delete(&order)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("order with ID %s has been deleted", id),
	})
}