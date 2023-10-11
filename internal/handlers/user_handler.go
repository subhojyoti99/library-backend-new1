package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"

	"library-backend-new1/internal/models"
)

// Create a new user
func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Create(&user)
		c.JSON(http.StatusCreated, user)
	}
}

// Update an existing user by ID
func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var user models.User

		if db.First(&user, id).RecordNotFound() {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Save(&user)
		c.JSON(http.StatusOK, user)
	}
}

// Get a user by ID
func GetUserByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var user models.User

		if db.First(&user, id).RecordNotFound() {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

// Delete a user by ID
func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var user models.User

		if db.First(&user, id).RecordNotFound() {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		db.Delete(&user)
		c.JSON(http.StatusNoContent, nil)
	}
}

// List all users
func ListUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.User
		db.Find(&users)
		c.JSON(http.StatusOK, users)
	}
}
