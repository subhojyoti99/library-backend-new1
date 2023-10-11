package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"

	"library-backend-new1/internal/models"
)

// Create a new library
func CreateLibrary(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var library models.Library
		err := c.BindJSON(&library)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the library already exists
		var existingLibrary models.Library
		if !db.Where("name = ?", library.Name).First(&existingLibrary).RecordNotFound() {
			c.JSON(http.StatusConflict, gin.H{"error": "Library already exists"})
			return
		}

		db.Create(&library)
		c.JSON(http.StatusCreated, library)
	}
}

// Update an existing library by ID
func UpdateLibrary(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var library models.Library

		if db.First(&library, id).RecordNotFound() {
			c.JSON(http.StatusNotFound, gin.H{"error": "Library not found"})
			return
		}

		err := c.BindJSON(&library)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Save(&library)
		c.JSON(http.StatusOK, library)
	}
}

// Get a library by ID
func GetLibraryByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var library models.Library

		if db.First(&library, id).RecordNotFound() {
			c.JSON(http.StatusNotFound, gin.H{"error": "Library not found"})
			return
		}

		c.JSON(http.StatusOK, library)
	}
}

// Delete a library by ID
func DeleteLibrary(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var library models.Library

		if db.First(&library, id).RecordNotFound() {
			c.JSON(http.StatusNotFound, gin.H{"error": "Library not found"})
			return
		}

		db.Delete(&library)
		c.JSON(http.StatusNoContent, nil)
	}
}

// List all libraries
func ListLibraries(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var libraries []models.Library
		db.Find(&libraries)
		c.JSON(http.StatusOK, libraries)
	}
}
