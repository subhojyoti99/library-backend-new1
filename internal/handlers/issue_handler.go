package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"

	"library-backend-new1/internal/models"
)

// Create a new issue registry
func CreateIssue(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var issue models.IssueRegistry
		err := c.BindJSON(&issue)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Create(&issue)
		c.JSON(http.StatusCreated, issue)
	}
}

// Update an existing issue registry by ID
func UpdateIssue(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var issue models.IssueRegistry

		if db.First(&issue, id).RecordNotFound() {
			c.JSON(http.StatusNotFound, gin.H{"error": "Issue registry not found"})
			return
		}

		err := c.BindJSON(&issue)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Save(&issue)
		c.JSON(http.StatusOK, issue)
	}
}

// Get an issue registry by ID
func GetIssueByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var issue models.IssueRegistry

		if db.First(&issue, id).RecordNotFound() {
			c.JSON(http.StatusNotFound, gin.H{"error": "Issue registry not found"})
			return
		}

		c.JSON(http.StatusOK, issue)
	}
}

// Delete an issue registry by ID
func DeleteIssue(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var issue models.IssueRegistry

		if db.First(&issue, id).RecordNotFound() {
			c.JSON(http.StatusNotFound, gin.H{"error": "Issue registry not found"})
			return
		}

		db.Delete(&issue)
		c.JSON(http.StatusNoContent, nil)
	}
}

// List all issue registries
func ListIssues(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var issues []models.IssueRegistry
		db.Find(&issues)
		c.JSON(http.StatusOK, issues)
	}
}
