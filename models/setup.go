package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("mysql", "root:welcome1@tcp(127.0.0.1:3306)/learngin")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Book{})
	DB = database
}
