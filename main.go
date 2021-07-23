package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"duguorong009/go-bookstore-rest-api-gin-gorm/models"

	"duguorong009/go-bookstore-rest-api-gin-gorm/controllers"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello guorong"})
	})

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)

	r.Run()
}
