package routes

import (
	"github.com/labstack/echo/v4"
	"myapp/controllers"
	"myapp/middlewares"
)

func SetupRouter(e *echo.Echo) {
	api := e.Group("/api/v1")
	{
		api.POST("/signup", controllers.Signup)

		login := api.Group("/login")
		{
			login.POST("", controllers.Login)

			// http://localhost:8001/api/v1/login/oauth
			login.GET("/oauth", controllers.OAuthStart)

			login.GET("/google", controllers.GoogleLogin)
			login.GET("/google/callback", controllers.GoogleLoginCallback)

			login.GET("/fb", controllers.FBLogin)
			login.GET("/fb/callback", controllers.FBLoginCallback)

			//login.GET("/twitter", controllers.FBLogin)
			//login.GET("/twitter/callback", controllers.FBLoginCallback)
		}

		authenticated := api.Group("")
		authenticated.Use(middlewares.Authz)
		{
			user := authenticated.Group("/user")
			{
				user.GET("/profile", controllers.UserProfile)
			}

			post := authenticated.Group("/post")
			{
				post.POST("", controllers.CreatePost)
				post.PUT("/:id", controllers.UpdatePost)
				post.DELETE("/:id", controllers.DeletePost)
			}
		}
	}
}
