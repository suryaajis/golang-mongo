package main

import (
	"golang-mongo/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	client, ctx, cancel, err := config.Connection("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer config.CloseConnection(client, ctx, cancel)
	config.Ping(client, ctx) // cek connection to database

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "welcome",
		})
	})

	router.Run()
}
