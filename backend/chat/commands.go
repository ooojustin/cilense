package chat

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func JoinRoom(data gin.H, ss *SocketSession, res *gin.H) {
	roomID := data["room_id"]
	ss.RoomID = roomID.(string)
	fmt.Println("user joining room", roomID)
	*res = map[string]interface{} {}
}

func SendMessage(data gin.H, ss *SocketSession, res *gin.H) {
	message := data["message"]
	fmt.Println("user sending message:", message)
	*res = map[string]interface{} {}
}
