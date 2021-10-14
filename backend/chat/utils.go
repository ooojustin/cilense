package chat

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func DoSendMessage(msg ChatMessage, ss *SocketSession) {

	// wait for 2 seconds (avoid concurrent write)
	time.Sleep(2 * time.Second)

	// new message data to send to clients
	msg.Sent = false
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
