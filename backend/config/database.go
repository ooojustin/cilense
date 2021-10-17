package config

import (
	"cilense.co/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := "cilense:Qpmpv$fWv3PR6P@tcp(db.cilense.co:3306)/cilense?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	database.AutoMigrate(&models.Room{})
	database.AutoMigrate(&models.Session{})
	DB = database
}
