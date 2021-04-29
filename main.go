package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"myapp/database"
	_ "myapp/docs"
	"myapp/helpers"
	"myapp/routes"
	"myapp/validators"
	"os"
)

// @title jph_app
// @version 1.0
// @description This is a pet project of written in echo
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@swagger.io

func main() {
	if helpers.LoadEnvVariables() != nil {
		log.Fatal("could not load env variables")
	}
	if database.InitDatabase() != nil {
		log.Fatal("could not connect to db")
	}

	e := echo.New()
	validators.InitValidator(e)
	routes.SetupRouter(e)

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}