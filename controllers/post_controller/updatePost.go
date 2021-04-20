package post_controller

import (
	"github.com/labstack/echo/v4"
	"myapp/helpers"
	"myapp/models"
	"net/http"
	"strconv"
)

// updates post in db
func UpdatePost(c echo.Context) error {
	postFromReq := new(models.Post)
	postIdFromReq, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res("postID was not in the url"))
	}

	if err := c.Bind(postFromReq); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res(err.Error()))
	}
	if err := c.Validate(postFromReq); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res(err.Error()))
	}

	// check if user is not updating somebody elses post
	userIdFromToken := c.Get("user_id").(int)
	postFromDB, err := postFromReq.GetPostByID(postIdFromReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res("could not find postFromReq in the DB"))
	}
	if userIdFromToken != postFromDB.UserID {
		return c.JSON(http.StatusBadRequest, helpers.Res("don't try to update post that you didn't create"))
	}

	postFromReq.UserID = userIdFromToken
	postFromReq.ID = postIdFromReq
	err = postFromReq.UpdatePostRecord()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Res(err.Error()))
	}

	return c.JSON(http.StatusAccepted, helpers.Res("post was updated"))
}
