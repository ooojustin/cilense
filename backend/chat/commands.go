package chat

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func JoinRoom(data gin.H, response *gin.H) {
	roomID := data["room_id"]
	fmt.Println("user joining room", roomID)
	*response = map[string]interface{} {}
}

func SendMessage(data gin.H, response *gin.H) {
	message := data["message"]
	fmt.Println("user sending message:", message)
	*response = map[string]interface{} {}
}
