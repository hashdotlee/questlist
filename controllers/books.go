package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"dblab/questlist/models"
	"dblab/questlist/initializers"
)

func FindBooks (c *gin.Context){
	var books []models.Book;
	initializers.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

type CreateBookInput struct {
  Title  string `json:"title" binding:"required"`
  Author string `json:"author" binding:"required"`
}

// POST /books
// Create new book
func CreateBook(c *gin.Context) {
  // Validate input
  var input CreateBookInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Create book
  book := models.Book{Title: input.Title, Author: input.Author}
  initializers.DB.Create(&book)

  c.JSON(http.StatusOK, gin.H{"data": book})
}

// GET /books/:id
// Find a book
func FindBook(c *gin.Context) {  // Get model if exist
  var book models.Book

  if err := initializers.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": book})
}

type UpdateBookInput struct {
  Title  string `json:"title"`
  Author string `json:"author"`  
}

// PATCH /books/:id
// Update a book
func UpdateBook(c *gin.Context) {
  // Get model if exist
  var book models.Book
  if err := initializers.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  // Validate input
  var input UpdateBookInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

    return
  }
  
  initializers.DB.Model(&book).Updates(input)

  c.JSON(http.StatusOK, gin.H{"data": book})
}

  // DELETE /books/:id
// Delete a book
func DeleteBook(c *gin.Context) {
  // Get model if exist
  var book models.Book
  if err := initializers.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  initializers.DB.Delete(&book)

  c.JSON(http.StatusOK, gin.H{"data": true})
}
