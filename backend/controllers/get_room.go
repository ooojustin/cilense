package controllers

import (
	"net/http"

	"cilense.co/config"
	"cilense.co/models"
	"github.com/gin-gonic/gin"
)

func GetRoomFromID(id string) *models.Room {
	// find room from database using id
	var room models.Room
	err := config.DB.Table("rooms").Where("id = ?", id).First(&room).Error
	if err == nil {
		return &room
	}
	return nil
}

func GetRoom(c *gin.Context) {
	// get the room from database
	id := c.Param("id")
	room := GetRoomFromID(id)
	if room == nil {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": room})
}
