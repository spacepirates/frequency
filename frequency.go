package main

import (
	"github.com/gin-gonic/gin"
	// "fmt"
	// "net/http"
	// "time"
)

func test(c *gin.Context) {
	c.JSON(200, gin.H{"test": "hello, web users. first change! oh yeah"})
}

func main() {
	g := gin.Default()

	//Routes
	g.GET("/", test)

	g.Run(":n00b")
}