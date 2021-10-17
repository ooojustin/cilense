package controllers

import (
	"net/http"

	"cilense.co/config"
	"cilense.co/models"
	"cilense.co/utils"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateRoom(c *gin.Context) {

	// initialize room and bind password
	room := models.Room{ID: uuid.NewV4()}
	c.BindJSON(&room)

	// hash password using bcrypt
	pwd := []byte(room.Password)
	hash, err := bcrypt.GenerateFromPassword(pwd, 14)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	room.Password = string(hash)

	// create session for owner containing the room and save to db
	// note that the room is created automatically with the session
	session := models.Session{
		ID:      uuid.NewV4(),
		Alias:   utils.GenerateRandomAlias(),
		Room:    &room,
		IsOwner: true,
	}
	config.DB.Create(&session)

	// return session object (which contains the room)
	c.JSON(http.StatusOK, gin.H{
		"session": session,
	})

}
