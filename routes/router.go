package routes

import (
	"github.com/labstack/echo/v4"
	"myapp/controllers/oauth_controller"
	"myapp/controllers/post_controller"
	"myapp/controllers/user_controller"
	"myapp/middlewares"
)

func SetupRouter(e *echo.Echo) {
	api := e.Group("/api/v1")
	{
		api.POST("/signup", user_controller.Signup)

		login := api.Group("/login")
		{
			login.POST("", user_controller.Login)

			// http://localhost:8001/api/v1/login/oauth
			login.GET("/oauth", oauth_controller.OAuthStart)

			login.GET("/google", oauth_controller.GoogleLogin)
			login.GET("/google/callback", oauth_controller.GoogleLoginCallback)

			login.GET("/fb", oauth_controller.FBLogin)
			login.GET("/fb/callback", oauth_controller.FBLoginCallback)
		}

		authenticated := api.Group("")
		authenticated.Use(middlewares.Authz)
		{
			user := authenticated.Group("/user")
			{
				user.GET("/profile", user_controller.UserProfile)
			}

			post := authenticated.Group("/post")
			{
				post.GET("/json", post_controller.GetAllPostsJSON)
				post.GET("/xml", post_controller.GetAllPostsXML)
				post.GET("/:id/json", post_controller.GetPostJSON)
				post.GET("/:id/xml", post_controller.GetPostXML)
				post.POST("", post_controller.CreatePost)
				post.PUT("/:id", post_controller.UpdatePost)
				post.DELETE("/:id", post_controller.DeletePost)
			}
		}
	}
}
