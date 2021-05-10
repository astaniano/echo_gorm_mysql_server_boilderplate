package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	_, err := hashPassword("secret")
	assert.NoError(t, err)
}

func TestCheckPassword(t *testing.T) {
	hashedPass, err := hashPassword("secret")
	assert.NoError(t, err)

	err = checkPassword(hashedPass, "secret")
	assert.NoError(t, err)
}
