package helpers

import "github.com/joho/godotenv"

func LoadEnvVariables() {
	err := godotenv.Load("/home/alex/my_files/PROGRAMMING/golang/server/.env")
	if err != nil {
		panic("could not load environmetal variables")
	}
}
