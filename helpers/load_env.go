package helpers

import "github.com/joho/godotenv"

func LoadEnvVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("could not load environmetal variables")
	}
}
