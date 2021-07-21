package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"./controllers"
	"./models"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello guorong"})
	})

	r.GET("/books", controllers.FindBooks)

	r.Run()
}
