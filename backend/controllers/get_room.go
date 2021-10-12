package controllers

import (
	"net/http"

	"cilense.co/config"
	"cilense.co/models"
	"github.com/gin-gonic/gin"
)

func GetRoom(c *gin.Context) {
	id := c.Param("id")
	var room models.Room
	config.DB.Table("rooms").Where("id = ?", id).First(&room)
	room.Password = "" // hide password from basic retrieval
	c.JSON(http.StatusOK, gin.H{"data": room})
}
