package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// TimeoutHandler adds a timeout to the request context
func TimeoutHandler(timeout time.Duration, next gin.HandlerFunc) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Create a context with timeout
        ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
        defer cancel()

        // Set the new context with timeout to the request
        c.Request = c.Request.WithContext(ctx)

        // Call the next handler (your logic)
        next(c)

        // Check if the context was canceled (timeout exceeded)
        if ctx.Err() == context.DeadlineExceeded {
            c.JSON(http.StatusGatewayTimeout, gin.H{
                "error": "request timed out",
            })
            c.Abort()
            return
        }
    }
}

// Simulates a long-running operation that respects context cancellation
func LongRunningOperation(c *gin.Context) {
    // Access the request context (which now has the timeout)
    ctx := c.Request.Context()

    // Simulating a long-running task (e.g., database query or external API call)
    select {
    case <-time.After(10 * time.Second): // Simulate work
        c.JSON(http.StatusOK, gin.H{
            "message": "Operation completed",
        })
    case <-ctx.Done(): // If the context is canceled (timeout exceeded)
        c.JSON(http.StatusRequestTimeout, gin.H{
            "error": "operation canceled",
        })
    }
}
