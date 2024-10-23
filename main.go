package main

import (
	"time"

	"github.com/cotopia-org/Event-Master/controllers"
	"github.com/cotopia-org/Event-Master/initializers"
	"github.com/cotopia-org/Event-Master/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToBD()
}

func main() {
	r := gin.Default()

	// Use the custom middleware
	r.Use(middlewares.TimeLogger())

	r.POST("/events", controllers.EventsCreate)
	r.GET("/events", controllers.EventsIndex)
	r.GET("/events/:id", controllers.EventsShow)
	r.PUT("/events/:id", controllers.EventsUpdate)
	r.DELETE("/events/:id", controllers.EventsDelete)

	// Route with timeout handling
	r.GET("/timeout", middlewares.TimeoutHandler(5*time.Second, middlewares.LongRunningOperation))

	// Define a simple GET route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:PORT
}