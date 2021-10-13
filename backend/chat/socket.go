package chat

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { 
		return true 
	},
}

type SocketAction struct {
	Action string `json:"action"`
	Data interface{} `json:"data"`
}

type SessionInfo struct {
	IsOwner bool `json:"is_owner"`
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	var si *SessionInfo
	si = new(SessionInfo)
	si.IsOwner = false

	for {

		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		// parse message from client
		var sa SocketAction
		err = json.Unmarshal(msg, &sa)
		if err != nil {
			break
		}
		data := sa.Data.(map[string]interface{})
		
		// execute action on client data to determine response
		var response gin.H
		switch sa.Action {
		case "join_room": JoinRoom(data, &response)
		case "send_message": SendMessage(data, &response)
		default:
			fmt.Println("Unhandled action:", sa.Action)
		}

		// serialize response to json and write to client
		rbytes, _ := json.Marshal(response)
		conn.WriteMessage(t, rbytes)
	
	}

}
