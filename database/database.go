package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitDatabase() (err error) {
	userAndPassword := os.Getenv("DB_USER") + ":" + os.Getenv("DB_USER_PASSWORD")
	DB, err = gorm.Open(mysql.Open(userAndPassword + "@tcp(127.0.0.1:3306)/jph?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}
