package database

import (
	"github.com/stretchr/testify/assert"
	"myapp/helpers"
	"testing"
)

func TestInitDB(t *testing.T) {
	err := helpers.LoadEnvVariables()
	assert.NoError(t, err)

	err = InitDatabase()
	assert.NoError(t, err)
}
