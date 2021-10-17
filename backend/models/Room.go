package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Room struct {
	ID        uuid.UUID `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Password  string    `json:"password,omitempty"`
}
