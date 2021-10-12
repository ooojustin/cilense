package main

import (
	"cilense.co/config"
	"cilense.co/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
	
func main() {


	config.InitDatabase()

    router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
        c.String(200, "gin!")
	})

	router.POST("/create_room", controllers.CreateRoom)

	router.Run()

}

