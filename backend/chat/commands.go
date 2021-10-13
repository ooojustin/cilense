package chat

import (
	"cilense.co/controllers"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func JoinRoom(data gin.H, res *gin.H, ss *SocketSession) {
	roomID := data["room_id"]
	ss.RoomID = roomID.(string)
	*res = gin.H{
		"session_id": ss.ID.String(),
	}
}

func SendMessage(data gin.H, res *gin.H, ss *SocketSession) {
	msg := ChatMessage{
		ID: uuid.NewV4(),
		Session: ss.ID,
		Text: data["message"].(string),
	}
	*res = gin.H{
		"data": msg,
	}
}

func Authenticate(data gin.H, res *gin.H, ss *SocketSession) {
	room := controllers.GetRoomFromID(ss.RoomID)
	if room.Token == data["token"] {
		ss.IsOwner = true
	}
	*res = gin.H{
		"success": ss.IsOwner,
	}
}
