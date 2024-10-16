package services

import (
	"net/http"
	"strconv"

	"api.com/models"
	"github.com/labstack/echo/v4"
)

func RegisterForEvent(c echo.Context) error {
	userId := c.Get("userId").(int64)
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Could not parse event id."})
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not fetch event."})
	}

	err = event.Register(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not register user for event."})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Registered!"})
}

func CancelRegistration(c echo.Context) error {
	userId := c.Get("userId").(int64)

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Could not parse event id."})
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not cancel registration."})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Cancelled!"})
}
