package main

import (
	"api.com/service"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", service.GetEvents)
	server.POST("/events", service.CreateEvent)

	server.Run(":8080")
}
