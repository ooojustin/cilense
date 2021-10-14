package chat

import (
	"cilense.co/models"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

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
	ID      uuid.UUID      `json:"id"`
	Alias   string         `json:"alias"`
	Text    string         `json:"text"`
	Sent    bool           `json:"sent"`
	Session *SocketSession `json:"-"`
}
