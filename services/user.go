package services

import (
	"fmt"
	"net/http"

	"api.com/models"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: cannot parse data": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save user"})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "user created", "user": user})
}
