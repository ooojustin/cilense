package models

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Code string
}
