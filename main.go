package main

import (
	"asignrest/database"
	"asignrest/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Koneksi()

	if err != nil {
		return 
	}
	r := gin.Default()

	routes.OrdersRoute(r,db)
	r.Run(":3000")

}