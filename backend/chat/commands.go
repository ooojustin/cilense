package chat

import (
	"cilense.co/config"
	"cilense.co/controllers"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func RestoreSession(data gin.H, res *gin.H, ss *SocketSession) {

	sessionID := data["session_id"].(string)
	session := controllers.GetSessionFromID(sessionID)

	if session != nil {
		ss.RoomID = session.RoomID
		ss.IsOwner = session.IsOwner
		ss.Model = session
		*res = gin.H{
			"type":    "session_restored",
			"session": session,
		}
	} else {
		*res = gin.H{
			"success": false,
			"message": "failed to restore session.",
		}
	}

}

func JoinRoom(data gin.H, res *gin.H, ss *SocketSession) {

	roomID := data["room_id"].(string)
	room := controllers.GetRoomFromID(roomID)

	ss.RoomID = roomID
	ss.Model.Room = room
	config.DB.Create(ss.Model)

	*res = gin.H{
		"type":       "room_joined",
		"session_id": ss.ID.String(),
	}

}

func SendMessage(data gin.H, res *gin.H, ss *SocketSession) {

	msg := ChatMessage{
		ID:      uuid.NewV4(),
		Session: ss.ID,
		Text:    data["message"].(string),
	}

	*res = gin.H{
		"data": msg,
	}

}

func Authenticate(data gin.H, res *gin.H, ss *SocketSession) {

	// use session to locate room from db
	room := controllers.GetRoomFromID(ss.RoomID)
	if room == nil {
		*res = gin.H{
			"success": false,
			"message": "Failed to authenticate - couldn't locate room.",
		}
		return
	}

	// set IsOwner to true
	if room.Token == data["token"] {
		ss.IsOwner = true
		if !ss.Model.IsOwner {
			ss.Model.IsOwner = true
			config.DB.Save(ss.Model)
		}
	}

	message := "Failed to authenticate."
	if ss.IsOwner {
		message = "Authentication successful."
	}
	*res = gin.H{
		"success": ss.IsOwner,
		"message": message,
	}

}
