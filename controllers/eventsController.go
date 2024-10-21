package controllers

import (
	"github.com/cotopia-org/Event-Master/initializers"
	"github.com/cotopia-org/Event-Master/models"
	"github.com/gin-gonic/gin"
)

func EventsCreate(c *gin.Context) {
	// get data off req body
	var body struct {
		Title string
		Body string
	}
	c.Bind(&body)

	// create event
	event := models.Event{Title: body.Title, Body: body.Body}
	
	result := initializers.DB.Create(&event)
	if result.Error != nil {
		c.Status(400)
		return
	}

	// return it
	c.JSON(201, gin.H{
		"event": event,
	})
}

func EventsIndex(c *gin.Context) {
	// get the events
	var events []models.Event
	initializers.DB.Find(&events)

	// respond with it
	c.JSON(200, gin.H{
		"event": events,
	})

}