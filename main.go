package main

import (
	"time"

	"github.com/cotopia-org/Event-Master/auth"
	"github.com/cotopia-org/Event-Master/controllers"
	docs "github.com/cotopia-org/Event-Master/docs"
	"github.com/cotopia-org/Event-Master/initializers"
	"github.com/cotopia-org/Event-Master/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

//	@title			EVENTMASTER API
//	@version		0.1
//	@description	It allows users to send various types of events which are stored, processed, and used to generate reports.

//	@contact.name	Ali Kharrati
//	@contact.email	ali.kharrati@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:3000
//	@BasePath	/

//	@securityDefinitions.basic	BasicAuth

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				JWT Bearer

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToBD()
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

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

	r.POST("/logic/intersect", controllers.IntersectSegments)
	r.POST("/logic/union", controllers.UnionSegments)
	r.POST("/logic/complement/:minBound/:maxBound", controllers.ComplementSegments)

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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:PORT
}
