package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitDatabase() (err error) {
	dbUserAndPassword := os.Getenv("DB_USER") + ":" + os.Getenv("DB_USER_PASSWORD")
	dbHostAndPort := os.Getenv("DB_HOST") + os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	DB, err = gorm.Open(mysql.Open(dbUserAndPassword+"@tcp("+dbHostAndPort+")/"+dbName+
		"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}
