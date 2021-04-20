package post_controller

import (
	"github.com/labstack/echo/v4"
	"myapp/database"
	"myapp/helpers"
	"myapp/models"
	"net/http"
	"strconv"
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

	return c.JSON(http.StatusCreated, helpers.Res("updated"))
}

// deletes post in db
func DeletePost(c echo.Context) error {
	postIdFromReq, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res("postID was not in the url"))
	}

	// check if user is not deleting somebody elses post
	userIdFromToken := c.Get("user_id").(int)
	PostFromDB := models.Post{}
	getPostRes := database.DB.First(&PostFromDB, postIdFromReq)
	if getPostRes.Error != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res("could not find post in the DB"))
	}
	if userIdFromToken != PostFromDB.UserID {
		return c.JSON(http.StatusBadRequest, helpers.Res("don't try to delete post that you didn't create"))
	}

	deletePostRes := database.DB.Delete(&models.Post{}, postIdFromReq)
	if deletePostRes.Error != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Res("could not delete post from the DB"))
	}

	return c.JSON(http.StatusCreated, helpers.Res("deleted"))
}
