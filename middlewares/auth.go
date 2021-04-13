package middlewares

import (
	"github.com/labstack/echo/v4"
	"myapp/auth"
	"myapp/helpers"
	"net/http"
	"os"
	"strings"
)

// Authz validates token and authorizes users
func Authz(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// clientToken := c.Request().Header.Get("Authorization")
		cookie, err := c.Cookie("Authorization")
		if err != nil {
			return err
		}

		clientToken := cookie.Value
		extractedToken := strings.Split(clientToken, "Bearer ")
		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			return c.JSON(http.StatusBadRequest, helpers.Res("Incorrect Format of Authorization Token"))
		}

		jwtWrapper := auth.JwtWrapper{
			SecretKey: os.Getenv("ACCESS_TOKEN_SECRET_KEY"),
			Issuer:    "AuthService",
		}
		claims, err := jwtWrapper.ValidateToken(clientToken)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helpers.Res(err.Error()))
		}

		c.Set("email", claims.Email)
		c.Set("user_id", claims.UserId)
		return next(c)
	}
}
