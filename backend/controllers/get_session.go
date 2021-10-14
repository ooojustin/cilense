package controllers

import (
	"cilense.co/config"
	"cilense.co/models"
)

func GetSessionFromID(id string) *models.Session {
	// find session from database using id
	var session models.Session
	err := config.DB.Table("sessions").Where("id = ?", id).First(&session).Error
	if err == nil {
		return &session
	}
	return nil
}
