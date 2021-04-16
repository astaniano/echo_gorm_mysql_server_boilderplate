package main

import (
	"myapp/database"
	"myapp/helpers"
	"myapp/models/models_post"
	"myapp/models/models_user"
)

func main() {
	helpers.LoadEnvVariables()
	database.InitDatabase()
	database.DB.AutoMigrate(&models_user.User{}, &models_post.Post{})
}