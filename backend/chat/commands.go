package chat

import (
	"cilense.co/config"
	"cilense.co/controllers"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func RestoreSession(data gin.H, res *gin.H, ss *SocketSession) {

	sessionID := data["session_id"].(string)
	session := controllers.GetSessionFromID(sessionID)

	if session != nil {
		ss.ID = session.ID
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

	// get room to join from database
	roomID := data["room_id"].(string)
	room := controllers.GetRoomFromID(roomID)

	// verify password
	hash := []byte(room.Password)
	pw := []byte(data["password"].(string))
	err := bcrypt.CompareHashAndPassword(hash, pw)
	if err != nil {
		*res = gin.H{
			"type": "wrong_password",
		}
		return
	}

	// create session model
	ss.RoomID = roomID
	ss.Model.Room = room
	config.DB.Create(ss.Model)

	// respond with session id for future use
	*res = gin.H{
		"type":       "room_joined",
		"session_id": ss.ID.String(),
	}

}

func SendMessage(data gin.H, res *gin.H, ss *SocketSession) {

	// initialize chat message
	msg := ChatMessage{
		ID:      uuid.NewV4(),
		Alias:   GenerateAlias(ss),
		Text:    data["message"].(string),
		Sent:    true,
		Session: ss,
		RoomID:  ss.RoomID,
	}

	// set response for message sender
	*res = gin.H{
		"type":    "message_sent",
		"message": msg,
	}

	// pass message to broadcast channel
	mchannel <- msg

}
