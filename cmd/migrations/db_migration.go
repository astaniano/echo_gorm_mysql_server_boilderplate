package main

import (
	"myapp/internal/app"
	"myapp/internal/models"
	"myapp/pkg/database"
)

func main() {
	err := app.InitApp()
	if err != nil {
		panic(err.Error())
	}

	db, err := database.InitDatabase()
	if err != nil {
		panic(err.Error())
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	err = db.AutoMigrate(
		&models.User{},
		&models.Post{},
	)
	if err != nil {
		panic(err.Error())
	}
}
