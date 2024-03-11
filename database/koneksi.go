package database

import (
	"asignrest/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Koneksi() (*gorm.DB, error) {
	dsn := "root:@tcp(localhost:3306)/go_rest_kominfo"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Items{}, &models.Orders{})
	if err != nil {
		return nil, err
	}
	DB = db
	fmt.Println("connecting to database")

	return db, nil
}