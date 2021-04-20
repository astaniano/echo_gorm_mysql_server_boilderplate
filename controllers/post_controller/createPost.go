package post_controller

import (
	"github.com/labstack/echo/v4"
	"myapp/helpers"
	"myapp/models"
	"net/http"
)

// creates a post in db
func CreatePost(c echo.Context) error {
	post := new(models.Post)

	if err := c.Bind(post); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res(err.Error()))
	}
	if err := c.Validate(post); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res(err.Error()))
	}

	post.UserID = c.Get("user_id").(int) // from the authorization middleware;
	err := post.CreatePostRecord()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Res(err.Error()))
	}

	return c.JSON(http.StatusCreated, helpers.Res("created"))
}
