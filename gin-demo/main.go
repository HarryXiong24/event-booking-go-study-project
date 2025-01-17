package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Msg struct {
	Status  int
	Message string
}

func main() {

	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.Header("Token", "jhghsg845jUIF83jh")
		c.Header("Content-Type", "application/text; charset=utf-8")
		c.JSON(http.StatusOK, &Msg{500, "ok"})
	})

	err := server.Run(":8080")
	if err != nil {
		return
	}
}
