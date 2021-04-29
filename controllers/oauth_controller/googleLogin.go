package oauth_controller

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"myapp/constants"
	"net/http"
	"os"
)

var googleOauthConfig *oauth2.Config

func GoogleLogin(c echo.Context) error {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  os.Getenv("HOST") + os.Getenv("PORT") + "/api/v1/login/google/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	url := googleOauthConfig.AuthCodeURL(constants.OAUTH_STATE_STRING)
	// url example: https://accounts.google.com/o/oauth2/auth?client_id=5555555&state=random11111
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleLoginCallback(c echo.Context) error {
	return oAuthCallbackHandler(
		c,
		"https://www.googleapis.com/oauth2/v2/userinfo?access_token=",
		googleOauthConfig,
	)
}
