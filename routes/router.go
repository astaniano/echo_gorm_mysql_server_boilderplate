package routes

import (
	"github.com/labstack/echo/v4"
	"myapp/controllers/controllers_post"
	"myapp/controllers/controllers_user"
	"myapp/middlewares"
)

func SetupRouter(e *echo.Echo) {
	api := e.Group("/api")
	{
		api.POST("/login", controllers_user.Login)
		api.POST("/signup", controllers_user.Signup)

		loggedIn := api.Group("")
		loggedIn.Use(middlewares.Authz)

		user := loggedIn.Group("/user")
		{
			user.GET("/profile", controllers_user.Profile)
		}

		post := loggedIn.Group("/post")
		{
			post.POST("", controllers_post.Create)
			post.PUT("", controllers_post.Update)
		}
	}
}
