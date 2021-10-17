package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Session struct {
	ID        uuid.UUID `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Alias     string    `json:"alias"`
	RoomID    string    `json:"-" gorm:"size:36"`
	Room      *Room     `json:"room" gorm:"foreignKey:RoomID;references:ID"`
	IsOwner   bool      `json:"is_owner"`
}
