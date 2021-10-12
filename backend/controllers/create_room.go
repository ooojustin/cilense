package controllers

import (
	"cilense.co/config"
	"cilense.co/models"
	"github.com/gin-gonic/gin"
)

func CreateRoom(c *gin.Context) {
	room := models.Room{Code: "whatever"}
	config.DB.Create(&room)
	c.JSON(200, gin.H{"data": room})
}
