package services

import (
	"fmt"
	"net/http"
	"strconv"

	"api.com/models"
	"github.com/labstack/echo/v4"
)

func GetEvent(c echo.Context) error {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not fetch events"})
		return err
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not fetch event"})
		return err
	}

	return c.JSON(http.StatusOK, event)
}

func GetEvents(c echo.Context) error {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not fetch events"})
		return err
	}

	return c.JSON(http.StatusOK, events)
}

func CreateEvent(c echo.Context) error {
	var event models.Event
	err := c.Bind(&event)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error: cannot parse data": err.Error()})
	}

	userId := c.Get("userId").(int64)
	event.UserID = userId

	err = event.Save()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not save event"})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"status": "event created", "event": event})
}

func UpdateEvent(c echo.Context) error {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not fetch events"})
		return err
	}

	userId := c.Get("userId").(int)
	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not fetch event"})
		return err
	}

	// Check if the user is the owner of the event
	if event.UserID != int64(userId) {
		c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorize to update event"})
		return err
	}

	var updateEvent models.Event
	err = c.Bind(&updateEvent) // This will bind the request body to the event struct
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error: cannot parse data": err.Error()})
		return err
	}

	updateEvent.ID = eventId

	err = updateEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not update event"})
		fmt.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"status": "event updated", "event": updateEvent})
}

func DeleteEvent(c echo.Context) error {
	// Code for deleting an event
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not fetch events"})
		return err
	}

	userId := c.Get("userId").(int)
	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not fetch event"})
		return err
	}

	if event.UserID != int64(userId) {
		c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorize to delete event"})
		return err
	}

	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not delete event"})
		fmt.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "event deleted"})
}
