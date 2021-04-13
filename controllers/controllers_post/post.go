package controllers_post

import (
	"github.com/labstack/echo/v4"
	"myapp/helpers"
	"myapp/models/models_post"
	"net/http"
)

// creates a post in db
func Create(c echo.Context) error {
	post := new(models_post.Post)

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

// updates a post in db
func Update(c echo.Context) error {
	postFromReq := new(models_post.Post)

	if err := c.Bind(postFromReq); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res(err.Error()))
	}
	if err := c.Validate(postFromReq); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res(err.Error()))
	}
	if postFromReq.ID == 0 {
		return c.JSON(http.StatusBadRequest, helpers.Res("postID was not in the req.body"))
	}

	// check if user is updating his own postFromReq
	userIdFromToken := c.Get("user_id").(int)
	postFromDB, err := postFromReq.GetPostByID(postFromReq.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res("could not find postFromReq in the DB"))
	}
	if userIdFromToken != postFromDB.UserID {
		return c.JSON(http.StatusBadRequest, helpers.Res("don't try to update postFromReq that you didn't create"))
	}

	postFromReq.UserID = userIdFromToken
	err = postFromReq.UpdatePostRecord()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Res(err.Error()))
	}

	return c.JSON(http.StatusCreated, "updated")
}