package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		fmt.Printf("%s %s - Started\n", c.Request.Method, c.Request.URL.Path)

		c.Next()

		fmt.Printf("%s - Completed in %v\n", c.Request.URL.Path, time.Since(start))
	}
}
