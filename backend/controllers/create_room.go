package controllers

import (
	"math/rand"
	"net/http"

	"cilense.co/config"
	"cilense.co/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateRoom(c *gin.Context) {

	// bind password to new room object
	room := models.Room{}
	c.BindJSON(&room)

	// hash password using bcrypt
	pwd := []byte(room.Password)
	hash, err := bcrypt.GenerateFromPassword(pwd, 14)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// store hashed password and random api token
	room.Password = string(hash)
	room.Token = RandomString(32)

	// submit query and return
	config.DB.Create(&room)
	c.JSON(http.StatusOK, gin.H{"data": room})

}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(n int) string {
	// generate a random string of a given length
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
