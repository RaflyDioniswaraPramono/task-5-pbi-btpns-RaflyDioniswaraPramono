package database

import (
	"example.com/go-rest-api-gin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/db_btpns_final_task?parseTime=true"))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Users{})
	db.AutoMigrate(&models.Photos{})

	DB = db
}
