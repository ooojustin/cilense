package controllers

import (
	"net/http"

	"cilense.co/config"
	"cilense.co/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateRoom(c *gin.Context) {

	room := models.Room{}
	config.DB.Create(&room)

	session := models.Session{
		ID:      uuid.NewV4(),
		Room:    &room,
		IsOwner: true,
	}
	c.BindJSON(&session)

	// hash password using bcrypt
	pwd := []byte(session.Password)
	hash, err := bcrypt.GenerateFromPassword(pwd, 14)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	session.Password = string(hash)
	config.DB.Create(&session)

	c.JSON(http.StatusOK, gin.H{
		"room":    room,
		"session": session,
	})

}
