package controllers_user

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"myapp/auth"
	"myapp/database"
	"myapp/helpers"
	"myapp/models/models_user"
	"net/http"
	"os"
	"time"
)

// Signup creates a controllers_user in db
func Signup(c echo.Context) error {
	user := new(models_user.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Error_res(err.Error()))
	}
	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Error_res(err.Error()))
	}

	err := user.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Error_res(err.Error()))
	}

	err = user.CreateUserRecord()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Error_res(err.Error()))
	}

	user.Password = ""
	return c.JSON(http.StatusCreated, user)
}

// LoginPayload login body
type LoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Login logs users in
func Login(c echo.Context) error {
	var payload LoginPayload
	var user models_user.User

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Error_res(err.Error()))
	}
	if err := c.Validate(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Error_res(err.Error()))
	}

	result := database.DB.Where("email = ?", payload.Email).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return c.JSON(http.StatusUnauthorized, helpers.Error_res("invalid user credentials"))
	}

	err := user.CheckPassword(payload.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helpers.Error_res(err.Error()))
	}

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       os.Getenv("TOKEN_SECRET_KEY"),
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Error_res(err.Error()))
	}

	cookie := &http.Cookie{
		Name:   "Authorization",
		Value:  "Bearer " + signedToken,
		HttpOnly: true, // disabling JavaScript access to cookie
		Expires: time.Now().Add(24 * time.Hour),
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, "ok")
}

// Profile returns controllers_user data
func Profile(c echo.Context) error {
	var user models_user.User

	email := c.Get("email") // from the authorization middleware

	result := database.DB.Where("email = ?", email.(string)).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		return c.JSON(http.StatusNotFound, helpers.Error_res("user not found"))
	}

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Error_res("could not get controllers_user profile"))
	}

	user.Password = ""
	return c.JSON(http.StatusOK, user)
}
