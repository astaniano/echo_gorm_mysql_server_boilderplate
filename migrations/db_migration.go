package main

import (
	"log"
	"myapp/database"
	"myapp/helpers"
	"myapp/models"
)

func main() {
	if helpers.LoadEnvVariables() != nil {
		log.Fatal("could not load env variables")
	}
	if database.InitDatabase() != nil {
		log.Fatal("could not connect to db")
	}

	err := database.DB.AutoMigrate(
		&models.User{},
		&models.Post{},
	)
	log.Fatal(err)
}
