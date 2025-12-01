package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello, World"})
	})

	if err := server.Run(":8000"); err != nil {
		log.Fatalf("Error starting the server: %v\n", err)
	}
}
