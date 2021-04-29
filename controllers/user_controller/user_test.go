package user_controller

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"myapp/database"
	"myapp/helpers"
	"myapp/models"
	"myapp/validators"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignUp(t *testing.T) {
	var actualResult models.User

	user := SignupPayload{
		FirstName: "Test User",
		LastName:  "Test User",
		Email:     "jwt@email.com",
		Password:  "secret",
	}

	payload, err := json.Marshal(&user)
	assert.NoError(t, err)

	e := echo.New()
	validators.InitValidator(e)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/signup", bytes.NewBuffer(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = helpers.LoadEnvVariables()
	assert.NoError(t, err)
	err = database.InitDatabase()
	assert.NoError(t, err)

	err = Signup(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		err = json.Unmarshal(rec.Body.Bytes(), &actualResult)
		assert.NoError(t, err)

		assert.Equal(t, user.FirstName, actualResult.FirstName)
		assert.Equal(t, user.Email, actualResult.Email)
	}

	database.DB.Unscoped().Delete(&actualResult)
}

func TestSignUpInvalidJSON(t *testing.T) {
	user := "test"

	payload, err := json.Marshal(&user)
	assert.NoError(t, err)

	e := echo.New()
	validators.InitValidator(e)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/signup", bytes.NewBuffer(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = Signup(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestLogin(t *testing.T) {
	err := helpers.LoadEnvVariables()
	assert.NoError(t, err)
	err = database.InitDatabase()
	assert.NoError(t, err)

	user := &models.User{
		FirstName: "Test User",
		LastName:  "Test User",
		Email:     "jwt@email.com",
		Password:  "secret",
	}
	err = user.HashPassword(user.Password)
	assert.NoError(t, err)
	err = user.CreateUserRecord()
	assert.NoError(t, err)

	controllersUser := LoginPayload{
		Email:    "jwt@email.com",
		Password: "secret",
	}
	payload, err := json.Marshal(&controllersUser)
	assert.NoError(t, err)

	e := echo.New()
	validators.InitValidator(e)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = Login(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	database.DB.Unscoped().Delete(&user)
}

func TestLoginInvalidJSON(t *testing.T) {
	controllersUser := "test"

	payload, err := json.Marshal(&controllersUser)
	assert.NoError(t, err)

	e := echo.New()
	validators.InitValidator(e)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = Login(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestLoginInvalidCredentials(t *testing.T) {
	controllersUser := LoginPayload{
		Email:    "jwt@email.com",
		Password: "invalid",
	}
	payload, err := json.Marshal(&controllersUser)
	assert.NoError(t, err)

	e := echo.New()
	validators.InitValidator(e)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = helpers.LoadEnvVariables()
	assert.NoError(t, err)
	err = database.InitDatabase()
	assert.NoError(t, err)

	err = Login(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestProfile(t *testing.T) {
	var profile models.User

	err := helpers.LoadEnvVariables()
	assert.NoError(t, err)
	err = database.InitDatabase()
	assert.NoError(t, err)

	user := models.User{
		FirstName: "Test User",
		LastName:  "Test User",
		Email:     "jwt@email.com",
		Password:  "secret",
	}
	err = user.HashPassword(user.Password)
	assert.NoError(t, err)
	err = user.CreateUserRecord()
	assert.NoError(t, err)

	e := echo.New()
	req, err := http.NewRequest(http.MethodGet, "/api/v1/user/profile", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("email", "jwt@email.com")

	err = UserProfile(c)
	assert.NoError(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &profile)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, user.Email, profile.Email)
	assert.Equal(t, user.FirstName, profile.FirstName)

	database.DB.Unscoped().Delete(&user)
}

func TestProfileNotFound(t *testing.T) {
	var profile models.User

	err := helpers.LoadEnvVariables()
	assert.NoError(t, err)
	err = database.InitDatabase()
	assert.NoError(t, err)

	e := echo.New()
	req, err := http.NewRequest(http.MethodGet, "/api/v1/user/profile", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("email", "nooooo@email.com")

	err = UserProfile(c)
	assert.NoError(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &profile)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)

	database.DB.Unscoped().Delete(&profile)
}
