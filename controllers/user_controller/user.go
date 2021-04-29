package user_controller

import (
	"github.com/labstack/echo/v4"
	"myapp/auth"
	"myapp/helpers"
	"myapp/models"
	"net/http"
	"os"
	"time"
)

// LoginPayload login body
type SignupPayload struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

// Signup godoc
// @Summary registers a new user
// @Description creates user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user_info body SignupPayload true "Sign up the user"
// @Success 201 {object} models.User
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/signup [post]
func Signup(c echo.Context) error {
	payload := new(SignupPayload)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res(err.Error()))
	}
	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res(err.Error()))
	}

	user := models.User{
		FirstName: payload.FirstName, LastName: payload.LastName, Email: payload.Email,
	}
	err := user.HashPassword(payload.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Res(err.Error()))
	}

	err = user.CreateUserRecord()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Res(err.Error()))
	}

	user.Password = ""
	return c.JSON(http.StatusCreated, &user)
}

// LoginPayload login body
type LoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Login logs users in
func Login(c echo.Context) error {
	var payload LoginPayload
	var user models.User

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res(err.Error()))
	}
	if err := c.Validate(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res(err.Error()))
	}

	err := user.FindUserByEmail(payload.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helpers.Res(err.Error()))
	}
	err = user.CheckPassword(payload.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helpers.Res(err.Error()))
	}

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       os.Getenv("ACCESS_TOKEN_SECRET_KEY"),
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}
	signedToken, err := jwtWrapper.GenerateToken(user.Email, user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Res(err.Error()))
	}

	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    "Bearer " + signedToken,
		HttpOnly: true, // disabling JavaScript access to cookie
		Expires:  time.Now().Add(24 * time.Hour),
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, "ok")
}

// UserProfile returns controllers_post data
func UserProfile(c echo.Context) error {
	user := new(models.User)

	email := c.Get("email") // from the authorization middleware

	err := user.FindUserByEmail(email.(string))
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.Res(err.Error()))
	}

	user.Password = ""
	return c.JSON(http.StatusOK, user)
}
