package helpers

import "github.com/joho/godotenv"

func LoadEnvVariables() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	return nil
}
