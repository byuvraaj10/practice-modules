package middleware

import (
	"database/sql"
	"encoding/base64"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(authHeader, "Basic "))
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid Authorization Header"})
			c.Abort()
			return
		}

		credentials := strings.SplitN(string(payload), ":", 2)
		if len(credentials) != 2 {
			c.JSON(401, gin.H{"error": "Invalid Credentials"})
			c.Abort()
			return
		}

		username, password := credentials[0], credentials[1]

		var storedPassword string
		query := "SELECT password FROM users WHERE username = ?"
		if err := db.QueryRow(query, username).Scan(&storedPassword); err != nil || storedPassword != password {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
