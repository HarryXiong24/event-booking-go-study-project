package services

import (
	"fmt"
	"net/http"
	"strconv"

	"api.com/models"
	"github.com/gin-gonic/gin"
)

func GetEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch events"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch event"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func GetEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch events"})
		return
	}

	c.JSON(http.StatusOK, events)

}

func CreateEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: cannot parse data": err.Error()})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save event"})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "event created", "event": event})
}

func UpdateEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch events"})
		return
	}

	_, err = models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch event"})
		return
	}

	var updateEvent models.Event
	err = c.ShouldBindJSON(&updateEvent) // This will bind the request body to the event struct
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: cannot parse data": err.Error()})
		return
	}

	updateEvent.ID = eventId

	err = updateEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update event"})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "event updated", "event": updateEvent})
}

func DeleteEvent(c *gin.Context) {
	// Code for deleting an event
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch events"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch event"})
		return
	}

	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete event"})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "event deleted"})
}
