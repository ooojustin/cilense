package controllers

import (
	"net/http"

	"cilense.co/config"
	"cilense.co/models"
	"github.com/gin-gonic/gin"
)

func GetRoom(c *gin.Context) {

	// find room from database using id
	id := c.Param("id")
	var room models.Room
	config.DB.Table("rooms").Where("id = ?", id).First(&room)

	// hide pw and api token from response
	room.Password = ""
	room.Token = ""

	c.JSON(http.StatusOK, gin.H{"data": room})

}
