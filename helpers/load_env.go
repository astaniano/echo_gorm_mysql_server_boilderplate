package helpers

import (
	"github.com/joho/godotenv"
	"path/filepath"
	"runtime"
)

func LoadEnvVariables() error {
	_, b, _, _ := runtime.Caller(0)
	pathToCurrentFile := filepath.Dir(b)

	err := godotenv.Load(pathToCurrentFile + "/../.env")
	if err != nil {
		return err
	}

	return nil
}
