package oauth_controller

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"myapp/constants"
	"net/http"
	"os"
)

var TwitterOauthConfig *oauth2.Config

func TwitterLogin(c echo.Context) error {
	// todo why &oauth2.Config ?????
	TwitterOauthConfig = &oauth2.Config{
		RedirectURL:  constants.HOST + constants.PORT + "/api/v1/login/fb/callback",
		ClientID:     os.Getenv("T_CLIENT_ID"),
		ClientSecret: os.Getenv("FB_CLIENT_SECRET"),
		Scopes:       []string{"public_profile", "email"},
		Endpoint:     facebook.Endpoint,
	}

	urlToRedirect := fbOauthConfig.AuthCodeURL(constants.OAUTH_STATE_STRING)
	return c.Redirect(http.StatusTemporaryRedirect, urlToRedirect)
}