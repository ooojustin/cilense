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

	// initialize room and bind password
	room := models.Room{}
	c.BindJSON(&room)

	// hash password using bcrypt
	pwd := []byte(room.Password)
	hash, err := bcrypt.GenerateFromPassword(pwd, 14)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// set hashed room password and save to db
	room.Password = string(hash)
	config.DB.Create(&room)

	// create session for owner and save to db
	session := models.Session{
		ID:      uuid.NewV4(),
		Room:    &room,
		IsOwner: true,
	}
	config.DB.Create(&session)

	// return session object (which contains the room)
	c.JSON(http.StatusOK, gin.H{
		"session": session,
	})

}
