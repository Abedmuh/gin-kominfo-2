package routes

import (
	"asignrest/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func OrdersRoute(r *gin.Engine, db *gorm.DB) {
	orderHandler:= controllers.NewOrderController(db)
	orderPath := r.Group("/order")

	orderPath.POST("/", orderHandler.CreateOrders)
	orderPath.GET("/", orderHandler.GetAllOrders)
	orderPath.PUT("/:id", orderHandler.UpdateOrders)
	orderPath.DELETE("/:id", orderHandler.DeleteOrder)
}