package helpers

import "github.com/joho/godotenv"

func LoadEnvVariables() error {
	err := godotenv.Load("/home/alex/my_files/PROGRAMMING/golang/server/.env")
	if err != nil {
		return err
	}

	return nil
}
