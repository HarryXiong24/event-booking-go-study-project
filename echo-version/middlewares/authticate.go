package middlewares

import (
	"net/http"

	"api.com/utils"
	"github.com/labstack/echo/v4"
)

func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		token := c.Request().Header.Get("Authorization")

		if token == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "no authorization"})
		}

		userId, err := utils.VerifyToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
		}

		c.Set("userId", userId) // Set the userId in the contexts

		return next(c) // Continue to the next middleware
	}
}
