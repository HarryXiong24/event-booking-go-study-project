package services

import (
	"fmt"
	"net/http"

	"api.com/models"
	"api.com/utils"
	"github.com/labstack/echo/v4"
)

func SignUp(c echo.Context) error {
	var user models.User

	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error: cannot parse data": err.Error()})
	}

	err = user.Save()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not save user"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"status": "user created"})
}

func Login(c echo.Context) error {
	var user models.User

	err := c.Bind(&user) // Bind the request body to the user struct
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error: cannot parse data": err.Error()})

	}

	err = user.ValidateCredentials()
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})

	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "logged in", "token": token})
}
