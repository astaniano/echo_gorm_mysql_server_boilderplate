package models

import (
	"github.com/stretchr/testify/assert"
	"myapp/database"
	"myapp/helpers"
	"testing"
)

func TestHashPassword(t *testing.T) {
	user := User{
		Password: "secret",
	}

	err := user.HashPassword(user.Password)
	assert.NoError(t, err)
}

func TestCreateUserRecord(t *testing.T) {
	var userResult User

	err := helpers.LoadEnvVariables()
	assert.NoError(t, err)
	err = database.InitDatabase()
	assert.NoError(t, err)

	user := User{
		FirstName: "Test User FN",
		LastName:  "Test User LN",
		Email:     "test@email.com",
		Password:  "secret",
	}
	err = user.HashPassword(user.Password)
	assert.NoError(t, err)
	err = user.CreateUserRecord()
	assert.NoError(t, err)

	database.DB.Where("email = ?", user.Email).Find(&userResult)
	assert.Equal(t, "Test User FN", userResult.FirstName)
	assert.Equal(t, "test@email.com", userResult.Email)

	database.DB.Unscoped().Delete(&user)
}

func TestCheckPassword(t *testing.T) {
	user := User{
		FirstName: "Test User FN",
		LastName:  "Test User LN",
		Email:     "test@email.com",
		Password:  "secret",
	}
	err := user.HashPassword(user.Password)
	assert.NoError(t, err)

	err = user.CheckPassword("secret")
	assert.NoError(t, err)
}
