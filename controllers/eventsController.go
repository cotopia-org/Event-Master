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

	// get the data off req body
	var body struct {
		Title string
		Body string
	}
	c.Bind(&body)

	// find the event we're updating
	var event models.Event
	initializers.DB.First(&event, id)

	// update it
	initializers.DB.Model(&event).Updates(models.Event{
		Title: body.Title,
		Body: body.Body,
	})

	// respond with it
	c.JSON(201, gin.H{
		"event": event,
	})
}

func EventsDelete(c *gin.Context) {
	// get the id off url
	id := c.Param("id")

	// delete the event
	initializers.DB.Delete(&models.Event{}, id)

	// respond!
	c.Status(200)

}