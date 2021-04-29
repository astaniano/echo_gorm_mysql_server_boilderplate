package routes

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"myapp/controllers/oauth_controller"
	"myapp/controllers/post_controller"
	"myapp/controllers/user_controller"
	"myapp/middlewares"
)

func SetupRouter(e *echo.Echo) {
	// http://localhost:8001/swagger/index.html
	e.GET("/swagger/*", echoSwagger.WrapHandler)

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
				post.GET("", post_controller.GetAllPosts)
				post.GET("/:id", post_controller.GetPost)
				post.POST("", post_controller.CreatePost)
				post.PUT("/:id", post_controller.UpdatePost)
				post.DELETE("/:id", post_controller.DeletePost)
			}
		}
	}
}

// CSRF CHECKING
//	e.GET("/csrf/check", func(c echo.Context) error {
//		resHtml := `<h1>hi</h1>
//<script>
//const response = fetch('http://localhost:8001/api/v1/login', {
//method: 'POST',
//headers: {
//  'Content-Type': 'application/json'
//},
//body: JSON.stringify({"email": "bro@ffffff.com", "password": "123456"})
//});
//</script>`
//		return c.HTML(200, resHtml)
//	})
