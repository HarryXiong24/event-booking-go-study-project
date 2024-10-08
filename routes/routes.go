package routes

import (
	"api.com/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Code for registering routes
	server.GET("/events", services.GetEvents)
	server.GET("/events/:id", services.GetEvent)
	server.POST("/events", services.CreateEvent)
	server.PUT("/events/:id", services.UpdateEvent)
	server.DELETE("/events/:id", services.DeleteEvent)
}
