package models

type Orders struct {
	Order_id      uint			`json:"id" gorm:"primaryKey;autoIncrement"`
	Ordered_at    string
	Customer_name string
}

type OrdersRequest struct {
	OrderedAt  string  `json:"orderedAt"`
	CustomName string 		`json:"customName"`
	Items      []struct {
		LineItem_id	uint				`json:"lineItem_id"`
		ItemCode    uint    		`json:"itemCode"`
		Description string 			`json:"description"`
		Quantity    uint    		`json:"quantity"`
	} `json:"items"`
}