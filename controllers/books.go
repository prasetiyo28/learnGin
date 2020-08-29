package controllers

import (
  "github.com/gin-gonic/gin"
  "github.com/prasetiyo28/learngin/models"
)

// GET /books
// Get all books
func FindBooks(c *gin.Context) {
  var books []models.Book
  models.DB.Find(&books)

  c.JSON(http.StatusOK, gin.H{"data": books})
}