package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Validates Content-Type header
func ValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Content-Type") != "application/json" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid content type"})
			c.Abort()
			return
		}
		c.Next()
	}
}
