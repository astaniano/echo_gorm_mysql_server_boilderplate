package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"myapp/database"
	_ "myapp/docs"
	"myapp/helpers"
	"myapp/routes"
	"myapp/validators"
)

// @title jph_app
// @version 1.0
// @description This is a pet project of written in echo
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@swagger.io

func main() {
	helpers.LoadEnvVariables()
	database.InitDatabase()

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	validators.InitValidator(e)
	routes.SetupRouter(e)

	e.Logger.Fatal(e.Start(":8001"))
}