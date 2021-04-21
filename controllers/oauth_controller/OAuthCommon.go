package oauth_controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"io/ioutil"
	"myapp/constants"
	"net/http"
	"net/url"
)

func OAuthStart(c echo.Context) error {
	var htmlIndex = `<html>
<body>
	<a href="/api/v1/login/google">Google Log In</a> <br/>
	<a href="/api/v1/login/fb">Facebook Log In</a> <br/>
</body>
</html>`

	return c.HTML(http.StatusOK, htmlIndex)
}

func oAuthCallbackHandler(c echo.Context, urlForGettingUserInfo string, config *oauth2.Config) error {
	content, err := getUserInfo(
		c.FormValue("state"),
		c.FormValue("code"),
		urlForGettingUserInfo,
		config,
	)
	if err != nil {
		fmt.Println(err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	//fmt.Fprintf(w, "Content: %s\n", content)
	return c.String(http.StatusOK, string(content))
}

// state: random1111111111
// code: 4/0AY0e-g........
func getUserInfo(state, code, urlForGettingUserInfo string, config *oauth2.Config) ([]byte, error) {
	if state != constants.OAUTH_STATE_STRING {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	// token.AccessToken: a0AfH6.......
	response, err := http.Get(urlForGettingUserInfo + url.QueryEscape(token.AccessToken))
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
