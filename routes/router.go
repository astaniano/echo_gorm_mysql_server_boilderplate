package routes

import (
	"github.com/labstack/echo/v4"
	"myapp/controllers/controllers_user"
	"myapp/controllers/controllers_post"
	"myapp/middlewares"
)

func SetupRouter(e *echo.Echo) {
	api := e.Group("/api")
	{
		api.POST("/login", controllers_user.Login)
		api.POST("/signup", controllers_user.Signup)

		user := api.Group("/user")
		user.Use(middlewares.Authz)
		{
			user.GET("/profile", controllers_user.Profile)
		}

		post := api.Group("/post")
		post.Use(middlewares.Authz)
		{
			post.POST("", controllers_post.Create)
			post.PUT("", controllers_post.Update)
		}
	}
}
