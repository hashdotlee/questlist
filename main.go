package main

import (
	_"net/http"
	"github.com/gin-gonic/gin"
	"dblab/questlist/models"
	"dblab/questlist/controllers"
)

func main() {
    r := gin.Default()
	
	// Initialize the database.
	models.ConnectDB()

	r.GET("/books", controllers.FindBooks)
	  r.POST("/books", controllers.CreateBook)
  r.GET("/books/:id", controllers.FindBook) // new
    r.PATCH("/books/:id", controllers.UpdateBook)
	  r.DELETE("/books/:id", controllers.DeleteBook)
	  
    r.Run("localhost:8080")
}

