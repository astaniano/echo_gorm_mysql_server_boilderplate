package auth

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	jwtWrapper := JwtWrapper{
		SecretKey:       "verySecret",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	generatedToken, err := jwtWrapper.GenerateToken("jwt@email.com", 1)
	assert.NoError(t, err)

	os.Setenv("testToken", generatedToken)
}

func TestValidateToken(t *testing.T) {
	encodedToken := os.Getenv("testToken")

	jwtWrapper := JwtWrapper{
		SecretKey: "verySecret",
		Issuer:    "AuthService",
	}

	claims, err := jwtWrapper.ValidateToken(encodedToken)
	assert.NoError(t, err)

	assert.Equal(t, "jwt@email.com", claims.Email)
	assert.Equal(t, "AuthService", claims.Issuer)
}
