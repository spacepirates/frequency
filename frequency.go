package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
	"math/rand"
)

func test(c *gin.Context) {
	name := ""

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
    num := r.Intn(2)

    switch num {
    	case 0:
	    	name = "mike"
	    	break;
	    case 1:
	    	name = "justin"
	    	break;
    }

	c.JSON(200, gin.H{"message of the day": name + " is a noob"})
}

func parseSound(c *gin.Context) {
	r := c.Request

    soundData := r.PostFormValue("sound_data")

    fmt.Println(soundData)

	c.JSON(200, gin.H{"sound_data": soundData})
}

func main() {
	g := gin.Default()

	//Routes
	g.GET("/", test)
	g.POST("/sound", parseSound)


	g.Run(":1337")

}