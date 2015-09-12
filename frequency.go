package main

import (
	"github.com/gin-gonic/gin"
	// "fmt"
	"net/http"
	// "time"
)

func main() {
	g := gin.Default()

	//Routes
	g.GET("/", c.JSON(200, gin.H{"test": "hello, web!"}))

	g.Run(":80")
}