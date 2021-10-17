package chat

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
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
				CloseSession(session)
			}

		}

	}

}

func AnnounceJoin(ss *SocketSession, re bool) {

	// joined or rejoined string
	wrd := "joined"
	if re {
		wrd = "re" + wrd
	}

	// create room message to announce user joined
	msg := ChatMessage{
		ID:            uuid.NewV4(),
		Text:          fmt.Sprintf("@%s has %s the room.", ss.Model.Alias, wrd),
		Session:       ss,
		RoomID:        ss.RoomID,
		IsRoomMessage: true,
	}

	// pass announcement to broadcast channel
	mchannel <- msg

}

func CloseSession(ss *SocketSession) {
	// close connection and remove session object from slice
	ss.Connection.Close()
	delete(sessions, ss)
}
