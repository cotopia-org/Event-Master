package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Custom middleware function to log request processing time
func TimeLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		startTime := time.Now()

		// Process request
		c.Next() // Next handler in the chain

		// Calculate and log the elapsed time
		elapsedTime := time.Since(startTime)
		fmt.Printf("-----> Request processed in %v\n", elapsedTime)
	}
}