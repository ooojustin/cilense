package chat

import (
	"encoding/json"
	"fmt"
	"net/http"

	"cilense.co/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

var clients = make(map[*websocket.Conn]bool)

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	} else {
		clients[conn] = true
	}

	// initialize session and store socket connection
	var ss *SocketSession
	ss = new(SocketSession)
	ss.ID = uuid.NewV4()
	ss.Connection = conn
	ss.Model = &models.Session{
		Room: nil,
	}
	ss.Model.ID = ss.ID

	// storing the connection like this allows us to send messages in other parts of the code
	// example: ss.Connection.WriteMessage(websocket.TextMessage, []byte("ping!"))

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
		case "restore_session":
			RestoreSession(data, &response, ss)
		case "join_room":
			JoinRoom(data, &response, ss)
		case "send_message":
			SendMessage(data, &response, ss)
		default:
			fmt.Println("Unhandled action:", sa.Action)
		}

		// serialize response to json and write to client
		// message types: https://github.com/gorilla/websocket/blob/master/conn.go#L62
		rbytes, _ := json.Marshal(response)
		conn.WriteMessage(t, rbytes)

	}

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type SocketAction struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

type SocketSession struct {
	ID         uuid.UUID       `json:"id"`
	Connection *websocket.Conn `json:"-"`
	RoomID     string          `json:"room_id"`
	IsOwner    bool            `json:"is_owner"`
	Model      *models.Session `json:"-"`
}

type ChatMessage struct {
	ID      uuid.UUID `json:"id"`
	Session uuid.UUID `json:"session"`
	Text    string    `json:"text"`
}
