package main

import (
	"api.com/db"
	"api.com/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	defer db.CloseDB()

	server := gin.Default()

	routers.RegisterRoutes(server)

	server.Run(":8081")
}
