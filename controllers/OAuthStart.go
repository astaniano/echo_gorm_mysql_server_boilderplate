package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func OAuthStart(c echo.Context) error {
	var htmlIndex = `<html>
<body>
	<a href="/api/v1/login/google">Google Log In</a> <br/>
	<a href="/api/v1/login/fb">Facebook Log In</a> <br/>
	<a href="/api/v1/login/twitter">Twitter Log In</a>
</body>
</html>`

	return c.HTML(http.StatusOK, htmlIndex)
}
