package main

import (
	"github.com/cotopia-org/Event-Master/controllers"
	"github.com/cotopia-org/Event-Master/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToBD()
}

func main() {
	r := gin.Default()

	r.POST("/events", controllers.EventsCreate)
	r.GET("/events", controllers.EventsIndex)
	r.GET("/events/:id", controllers.EventsShow)
	r.PUT("/events/:id", controllers.EventsUpdate)

	r.Run() // listen and serve on 0.0.0.0:PORT
}