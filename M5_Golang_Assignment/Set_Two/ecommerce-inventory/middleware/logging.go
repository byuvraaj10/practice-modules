package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Logs request method, path, status, and duration
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		fmt.Printf("%s %s | Status: %d | Time: %v\n", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), time.Since(startTime))
	}
}
