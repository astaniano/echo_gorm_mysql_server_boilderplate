package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"io/ioutil"
	"myapp/constants"
	"net/url"
	"os"
	"net/http"
)

var fbOauthConfig *oauth2.Config

func FBLogin(c echo.Context) error {
	// todo why &oauth2.Config ?????
	fbOauthConfig = &oauth2.Config{
		RedirectURL:  constants.HOST + constants.PORT + "/api/v1/login/fb/callback",
		ClientID:     os.Getenv("FB_CLIENT_ID"),
		ClientSecret: os.Getenv("FB_CLIENT_SECRET"),
		Scopes:   []string{"public_profile", "email"},
		Endpoint: facebook.Endpoint,
	}

	urlToRedirect := fbOauthConfig.AuthCodeURL(constants.OAUTH_STATE_STRING)
	return c.Redirect(http.StatusTemporaryRedirect, urlToRedirect)
}

func FBLoginCallback(c echo.Context) error {
	content, err := getUserInfoFB(c.FormValue("state"), c.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	return c.String(http.StatusOK, string(content))
}

func getUserInfoFB(state string, code string) ([]byte, error) {
	if state != constants.OAUTH_STATE_STRING {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := fbOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	// token.AccessToken: a0AfH6.......
	response, err := http.Get("https://graph.facebook.com/me?access_token=" +
		url.QueryEscape(token.AccessToken))
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