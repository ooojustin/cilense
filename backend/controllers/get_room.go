package controllers

import (
	"cilense.co/config"
	"cilense.co/models"
	"github.com/gin-gonic/gin"
)

func GetRoom(c *gin.Context) {
	id := c.Param("id")
	var r models.Room
	config.DB.Table("rooms").Where("id = ?", id).First(&r)
	c.JSON(200, gin.H{"data": r})
}
