package main

import (
	"github.com/labstack/echo/v4"
	"myapp/database"
	"myapp/helpers"
	"myapp/routes"
	"myapp/validators"
)

func main() {
	err := database.InitDatabase()
	if err != nil {
		panic("could not open database connection")
	}

	helpers.LoadEnvVariables()

	e := echo.New()
	validators.InitValidator(e)
	routes.SetupRouter(e)

	e.Logger.Fatal(e.Start(":8001"))
}