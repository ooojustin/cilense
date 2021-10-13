package chat

import (
	"fmt"

	"cilense.co/controllers"
	"github.com/gin-gonic/gin"
)

func JoinRoom(data gin.H, res *gin.H, ss *SocketSession) {
	roomID := data["room_id"]
	ss.RoomID = roomID.(string)
	fmt.Println("user joining room", roomID)
	*res = gin.H {}
}

func SendMessage(data gin.H, res *gin.H, ss *SocketSession) {
	message := data["message"]
	fmt.Println("user sending message:", message)
	*res = gin.H {}
}

func Authenticate(data gin.H, res *gin.H, ss *SocketSession) {
	room := controllers.GetRoomFromID(ss.RoomID)
	if room.Token == data["token"] {
		ss.IsOwner = true
	}
	*res = gin.H {
		"success": ss.IsOwner,
	}
}
