package controllers

import (
	"net/http"

	"cilense.co/config"
	"cilense.co/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateRoom(c *gin.Context) {
	room := models.Room{}
	c.BindJSON(&room)
	pwd := []byte(room.Password)
	hash, err := bcrypt.GenerateFromPassword(pwd, 14)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	room.Password = string(hash)
	config.DB.Create(&room)
	c.JSON(http.StatusOK, gin.H{"data": room})
}
