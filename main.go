package main

import (
	"api.com/db"
	"api.com/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	defer db.CloseDB()

	server := gin.Default()

	server.GET("/events", service.GetEvents)
	server.POST("/events", service.CreateEvent)

	server.Run(":8080")
}
