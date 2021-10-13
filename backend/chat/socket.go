package chat

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {

		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var sa SocketAction
		err = json.Unmarshal(msg, &sa)
		if err != nil {
			break
		}
	
		data := sa.Data.(map[string]interface{})
		switch sa.Action {
		case "join_room":
			roomID := data["room_id"]
			fmt.Println("user joining room", roomID)
		default:
			fmt.Println("Unhandled action:", sa.Action)
		}


		conn.WriteMessage(t, msg)
	
	}

}
