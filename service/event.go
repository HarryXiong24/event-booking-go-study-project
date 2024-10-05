package service

import (
	"net/http"

	"api.com/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(http.StatusOK, events)

}

func CreateEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: cannot parse data": err.Error()})
		return
	}

	event.ID = len(models.GetAllEvents()) + 1
	event.UserID = 1

	event.Save()
	c.JSON(http.StatusCreated, gin.H{"status": "event created", "event": event})
}
