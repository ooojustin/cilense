package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Session struct {
	ID        uuid.UUID `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	RoomID    string    `json:"-"`
	Room      *Room     `json:"room" gorm:"foreignKey:RoomID;references:ID"`
	IsOwner   bool      `json:"is_owner"`
}