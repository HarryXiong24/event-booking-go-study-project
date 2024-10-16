package routers

import (
	"api.com/middlewares"
	"api.com/services"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	// Code for registering routes
	e.GET("/events", services.GetEvents)
	e.GET("/events/:id", services.GetEvent)

	authenticated := e.Group("")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", services.CreateEvent)
	authenticated.PUT("/events/:id", services.UpdateEvent)
	authenticated.DELETE("/events/:id", services.DeleteEvent)

	authenticated.POST("/events/:id/register", services.RegisterForEvent)
	authenticated.DELETE("/events/:id/register", services.CancelRegistration)

	e.POST("/sign-up", services.SignUp)
	e.POST("/login", services.Login)
}
