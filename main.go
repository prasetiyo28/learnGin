package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prasetiyo28/learnGin/controllers"
	"github.com/prasetiyo28/learnGin/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	// })
	r.GET("/books", controllers.FindBooks) // new
	r.Run()
}
