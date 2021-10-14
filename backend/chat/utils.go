package chat

import (
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func MessageSender() {

	for {

		// read new message from the message channel
		msg := <-mchannel

		// new message data to send to clients
		msg.Sent = false
		packet := gin.H{
			"type":    "new_message",
			"message": msg,
		}
		pbytes, _ := json.Marshal(packet)

		for session := range sessions {

			// don't send this to the message sender
			if session == msg.Session {
				continue
			}

			// ensure we only send it to people in the correct room
			if session.RoomID != msg.RoomID {
				continue
			}

			// send message to client / handle failure
			err := session.Connection.WriteMessage(websocket.TextMessage, pbytes)
			if err != nil {
				session.Connection.Close()
				delete(sessions, session)
			}

		}

	}

}

func GenerateAlias(ss *SocketSession) string {
	alias := ""
	slice := strings.Split(ss.ID.String(), "-")
	for i := range slice {
		item := slice[i]
		alias += item[0:1]
		alias += item[len(item)-1:]
	}
	return alias
}
