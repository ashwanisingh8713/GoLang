package controllers

import (
	models2 "book/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// FindBooks godoc
// @Summary Finds all books
// @Description To get all books
// @Tags Create & Get Book
// @Produce json
// @Success 200 {object} models.Book
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /books [get]
func FindBooks(c *gin.Context) {
	var books []models2.Book
	models2.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// GET /books/:id
// Find a book

// FindBook godoc
// @Summary Find a book based on id
// @Description To get all books
// @Tags Create & Get Book
// @Produce json
// @Success 200 {object} models.Book
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /books/:id [get]
func FindBook(c *gin.Context) {
	// Get model if exist
	var book models2.Book
	if err := models2.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// CreateBook godoc
// @Summary Create book
// @Description To create a new book
// @Tags Create & Get Book
// @Produce json
// @Success 200 {object} models.Book
// @Param       json  body models.Book true "It takes book infos"
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /books [post]
func CreateBook(c *gin.Context) {
	// Validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models2.Book{Title: input.Title, Author: input.Author}
	models2.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook godoc
// @Summary Update book
// @Description To update exist book
// @Tags Update & Delete Book
// @Produce json
// @Success 200 {object} models.Book
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /books/:id [patch]
func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models2.Book
	if err := models2.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models2.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook godoc
// @Summary Delete book
// @Description To delete exist book
// @Tags Update & Delete Book
// @Produce json
// @Success 200 {object} models.Book
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /books/:id [delete]
func DeleteBook(c *gin.Context) {
	// Get model if exist
	var book models2.Book
	if err := models2.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models2.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
