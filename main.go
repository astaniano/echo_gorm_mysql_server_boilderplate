package main

import (
	"github.com/labstack/echo/v4"
	"myapp/database"
	"myapp/helpers"
	"myapp/routes"
	"myapp/validators"
)

func main() {
	helpers.LoadEnvVariables()
	database.InitDatabase()

	e := echo.New()
	validators.InitValidator(e)
	routes.SetupRouter(e)

	e.Logger.Fatal(e.Start(":8001"))
}