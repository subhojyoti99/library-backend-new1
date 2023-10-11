package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"

	"library-backend-new1/internal/models"
)

// Create a new book
func CreateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.BookInventory
		err := c.BindJSON(&book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Create(&book)
		c.JSON(http.StatusCreated, book)
	}
}

// Update an existing book by ISBN
func UpdateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		isbn := c.Param("isbn")
		var book models.BookInventory

		if db.First(&book, "isbn = ?", isbn).RecordNotFound() {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		err := c.BindJSON(&book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Save(&book)
		c.JSON(http.StatusOK, book)
	}
}

// Get a book by ISBN
func GetBookByISBN(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		isbn := c.Param("isbn")
		var book models.BookInventory

		if db.First(&book, "isbn = ?", isbn).RecordNotFound() {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		c.JSON(http.StatusOK, book)
	}
}

// Delete a book by ISBN
func DeleteBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		isbn := c.Param("isbn")
		var book models.BookInventory

		if db.First(&book, "isbn = ?", isbn).RecordNotFound() {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		db.Delete(&book)
		c.JSON(http.StatusNoContent, nil)
	}
}

// List all books
func ListBooks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var books []models.BookInventory
		db.Find(&books)
		c.JSON(http.StatusOK, books)
	}
}
