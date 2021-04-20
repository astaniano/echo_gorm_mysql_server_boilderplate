package main

import (
	"myapp/database"
	"myapp/helpers"
	"myapp/models"
)

func main() {
	helpers.LoadEnvVariables()
	database.InitDatabase()
	database.DB.AutoMigrate(&models.User{}, &models.Post{})
}