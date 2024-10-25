package main

import (
	"time"

	"github.com/cotopia-org/Event-Master/auth"
	"github.com/cotopia-org/Event-Master/controllers"
	"github.com/cotopia-org/Event-Master/initializers"
	"github.com/cotopia-org/Event-Master/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToBD()
}

func main() {
	r := gin.Default()

	// CORS middleware configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Use the custom middleware
	r.Use(middlewares.TimeLogger())

	r.POST("/users", controllers.UserCreate)
	r.GET("/users/:id", controllers.UserGetByID)
	r.PUT("/users/:id", controllers.UserUpdate)
	r.DELETE("/users/:id", controllers.UserDelete)

	r.POST("/login", auth.Login)

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

	// Protected route group
	protected := r.Group("/protected")
	protected.Use(auth.JWTAuthMiddleware())
	{
		protected.GET("/dashboard", auth.ProtectedRoute)
	}

	r.Run() // listen and serve on 0.0.0.0:PORT
}
