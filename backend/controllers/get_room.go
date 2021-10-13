package controllers

import (
	"net/http"

	"cilense.co/config"
	"cilense.co/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetRoomFromID(id string) models.Room {
	// find room from database using id
	var room models.Room
	config.DB.Table("rooms").Where("id = ?", id).First(&room)
	return room
}

func GetRoom(c *gin.Context) {

	// get the room from database
	id := c.Param("id")
	room := GetRoomFromID(id)

	// hide pw and api token from response
	room.Password = ""
	room.Token = ""

	c.JSON(http.StatusOK, gin.H{"data": room})

}

func GetRoomAuthenticated(c *gin.Context) {

	// get the rooom from database
	id := c.Param("id")
	room := GetRoomFromID(id)

	// compare password in request to hashed password
	hash := room.Password
	c.BindJSON(&room)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(room.Password))

	// if the password was correct, return the whole room object (with token)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": room,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error": "Failed to authorize request.",
		})
	}

}
