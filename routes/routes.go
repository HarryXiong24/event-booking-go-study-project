package routes

import (
	"api.com/middlewares"
	"api.com/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Code for registering routes
	server.GET("/events", services.GetEvents)
	server.GET("/events/:id", services.GetEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", services.CreateEvent)
	authenticated.PUT("/events/:id", services.UpdateEvent)
	authenticated.DELETE("/events/:id", services.DeleteEvent)

	server.POST("/sign-up", services.SignUp)
	server.POST("/login", services.Login)
}
