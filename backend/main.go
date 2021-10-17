package main

import (
	"math/rand"
	"net/http"
	"time"

	"cilense.co/chat"
	"cilense.co/config"
	"cilense.co/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// initialize psuedo-random library with seed
	rand.Seed(time.Now().UnixNano())

	// connect to database and run goroutine to process message channel
	config.InitDatabase()
	go chat.MessageSender()

	// configure gin router with cors middleware
	router := gin.Default()
	router.Use(cors.Default())

	// backend index and new websocket connection handler
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "cilense.co")
	})
	router.GET("/ws", func(c *gin.Context) {
		chat.WebSocketHandler(c.Writer, c.Request)
	})

	// room model endpoints
	router.POST("/create_room", controllers.CreateRoom)
	router.GET("/room/:id", controllers.GetRoom)

	router.Run()

}
