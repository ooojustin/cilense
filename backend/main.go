package main

import (
	"cilense.co/config"
	"cilense.co/controllers"
	"github.com/gin-gonic/gin"
)
	
func main() {


	config.InitDatabase()

    r := gin.Default()

	r.GET("/", func(c *gin.Context) {
        c.String(200, "gin!")
	})

	r.POST("/create_room", controllers.CreateRoom)

	r.Run()

}

