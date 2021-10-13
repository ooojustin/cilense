package controllers

import (
	"net/http"

	"cilense.co/config"
	"cilense.co/models"
	"github.com/gin-gonic/gin"
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
