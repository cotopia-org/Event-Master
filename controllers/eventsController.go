package controllers

import (
	"net/http"

	"github.com/cotopia-org/Event-Master/initializers"
	"github.com/cotopia-org/Event-Master/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EventsCreate(c *gin.Context) {
	// get data off req body
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create event
	if err := initializers.DB.Create(&event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return it
	c.JSON(http.StatusCreated, event)
}

func EventsIndex(c *gin.Context) {
	// get the events
	var events []models.Event
	initializers.DB.Find(&events)

	// respond with it
	c.JSON(200, gin.H{
		"events": events,
	})

}

func EventsShow(c *gin.Context) {
	// get id off url
	id := c.Param("id")

	// get the events
	var event models.Event
	initializers.DB.First(&event, id)

	// respond with it
	c.JSON(200, gin.H{
		"event": event,
	})
}

func EventsUpdate(c *gin.Context) {
	// get the id off url
	id := c.Param("id")

	// find the event we're updating
	var event models.Event

	if err := initializers.DB.First(&event, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// get the data off req body
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update it
	if err := initializers.DB.Save(&event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// respond with it
	c.JSON(http.StatusOK, event)
}

func EventsDelete(c *gin.Context) {
	// get the id off url
	id := c.Param("id")

	// delete the event
	initializers.DB.Delete(&models.Event{}, id)

	// respond!
	c.Status(200)

}
