package database

import (
	"github.com/stretchr/testify/assert"
	"myapp/helpers"
	"testing"
)

func TestInitDB(t *testing.T) {
	helpers.LoadEnvVariables()
	err := InitDatabase()
	assert.NoError(t, err)
}
