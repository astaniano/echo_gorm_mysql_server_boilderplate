package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() (err error) {
	DB, err = gorm.Open(mysql.Open("sammy:sammy@tcp(127.0.0.1:3306)/jph?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return
	}

	return
}
