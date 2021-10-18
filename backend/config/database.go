package config

import (
	"cilense.co/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {

	// dynamically select between databases based on release mode
	// if in release mode, use mysql database connection - otherwise sqlite is used with a local db file
	var err error
	if gin.Mode() == gin.ReleaseMode {
		dsn := "cilense:Qpmpv$fWv3PR6P@tcp(db.cilense.co:3306)/cilense?charset=utf8mb4&parseTime=True&loc=Local"
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else {
		DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	}

	// panic if there's an error
	if err != nil {
		panic("failed to connect database")
	}

	// run model migrations
	DB.AutoMigrate(&models.Room{})
	DB.AutoMigrate(&models.Session{})

}
