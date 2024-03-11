package models

type Items struct {
	Item_id     uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Item_code   uint   `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
}