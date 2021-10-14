package chat

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func DoSendMessage(msg ChatMessage, ss *SocketSession) {

	// pause for 5 seconds (avoid concurrent write)
	time.Sleep(5 * time.Second)

	// new message data to send to clients
	packet := gin.H{
		"type":    "new_message",
		"message": msg,
	}
	pbytes, _ := json.Marshal(packet)

	for client := range clients {

		// don't send message to sender (it's been stored already)
		if client == ss.Connection {
			continue
		}

		// send message to client / handle failure
		err := client.WriteMessage(websocket.TextMessage, pbytes)
		if err != nil {
			client.Close()
			delete(clients, client)
		}

	}

}
