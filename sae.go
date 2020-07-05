package main

import (
	"io"
	"os"

	gin "github.com/gin-gonic/gin"
)

func main() {

	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("./logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()

	router.StaticFile("/", "./public/index.html")
	router.StaticFile("/favicon.png", "./public/favicon.png")

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong!",
		})
	})

	router.Run(":7200")
}
