package middleware

import (
	"library-backend-new1/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get email from request headers
		email := c.GetHeader("Email")

		// Check if the user exists
		var user models.User
		if db.Where("email = ?", email).First(&user).RecordNotFound() {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Set user information in the context for later use
		c.Set("user", user)

		// Continue processing the request
		c.Next()
	}
}
