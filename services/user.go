package services

import (
	"fmt"
	"net/http"

	"api.com/models"
	"api.com/utils"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: cannot parse data": err.Error()})
		fmt.Println(err)
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save user"})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "user created"})
}

func Login(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user) // Bind the request body to the user struct
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: cannot parse data": err.Error()})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "logged in", "token": token})
}
