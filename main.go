package main

import (
	"api.com/db"
	"api.com/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	defer db.CloseDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
