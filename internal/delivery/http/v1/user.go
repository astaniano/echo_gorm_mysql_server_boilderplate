package v1

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/helpers"
	"myapp/internal/shared/payloads"
	"net/http"
	"time"
)

// Signup godoc
// @Summary registers a new user
// @Description creates user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user_info body payloads.SignUpPayload true "Sign up the user"
// @Success 201 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/signup [post]
func (h *Handler) signUp(c echo.Context) error {
	payload := new(payloads.SignUpPayload)
	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if  err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	err := h.services.User.SignUp(payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, helpers.Res("user was created"))
}

// Login logs users in
func (h *Handler) signIn(c echo.Context) error {
	payload := new(payloads.SignInPayload)
	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, helpers.Res(err.Error()))
	}
	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, helpers.Res(err.Error()))
	}

	signedToken, err := h.services.User.SignIn(payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, helpers.Res(err.Error()))
	}

	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    "Bearer " + signedToken,
		HttpOnly: true, // disabling JavaScript access to cookie
		Expires:  time.Now().Add(24 * time.Hour),
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, helpers.Res("ok"))
}

func (h *Handler) getUserProfile(c echo.Context) error {
	email := c.Get("email") // from the authorization middleware
	user, err := h.services.User.GetUserProfile(email.(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, helpers.Res(err.Error()))
	}

	user.Password = ""
	return c.JSON(http.StatusOK, user)
}
