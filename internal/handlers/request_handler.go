package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"

	"library-backend-new1/internal/models"
)

// Create a new request event
func CreateRequest(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request models.RequestEvents
		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Create(&request)
		c.JSON(http.StatusCreated, request)
	}
}

// Update an existing request event by ID
func UpdateRequest(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var request models.RequestEvents

		if db.First(&request, id).RecordNotFound() {
			c.JSON(http.StatusNotFound, gin.H{"error": "Request event not found"})
			return
		}

		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Save(&request)
		c.JSON(http.StatusOK, request)
	}
}

// Get a request event by ID
func GetRequestByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var request models.RequestEvents

		if db.First(&request, id).RecordNotFound() {
			c.JSON(http.StatusNotFound, gin.H{"error": "Request event not found"})
			return
		}

		c.JSON(http.StatusOK, request)
	}
}

// Delete a request event by ID
func DeleteRequest(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var request models.RequestEvents

		if db.First(&request, id).RecordNotFound() {
			c.JSON(http.StatusNotFound, gin.H{"error": "Request event not found"})
			return
		}

		db.Delete(&request)
		c.JSON(http.StatusNoContent, nil)
	}
}

// List all request events
func ListRequests(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requests []models.RequestEvents
		db.Find(&requests)
		c.JSON(http.StatusOK, requests)
	}
}
