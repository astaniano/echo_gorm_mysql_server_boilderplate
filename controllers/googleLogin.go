package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"myapp/constants"
	"os"
	"io/ioutil"
	"net/http"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig oauth2.Config

func GoogleLogin(c echo.Context) error {
	googleOauthConfig = oauth2.Config{
		RedirectURL:  constants.HOST + constants.PORT + "/api/v1/login/google/callback",
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
	content, err := getUserInfo(c.FormValue("state"), c.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	//fmt.Fprintf(w, "Content: %s\n", content)
	return c.String(http.StatusOK, string(content))
}

/**
@param state: random1111111111
@param code: 4/0AY0e-g........

@return Content: {
  "id": "5555555555",
  "email": "example@gmail.com",
  "verified_email": true,
  "picture": "https://........."
}
*/
func getUserInfo(state string, code string) ([]byte, error) {
	if state != constants.OAUTH_STATE_STRING {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	// token.AccessToken: a0AfH6.......
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, nil
}
