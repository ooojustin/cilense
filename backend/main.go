package main

import (
	"net/http"

	"cilense.co/chat"
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
        c.String(http.StatusOK, "gin!")
	})

	router.GET("/ws", func(c *gin.Context) {
		chat.WebSocketHandler(c.Writer, c.Request)
	})

	router.POST("/create_room", controllers.CreateRoom)
	router.POST("/room/:id", controllers.GetRoomAuthenticated)
	router.GET("/room/:id", controllers.GetRoom)

	router.Run()

}

