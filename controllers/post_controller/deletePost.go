package post_controller

import (
	"github.com/labstack/echo/v4"
	"myapp/database"
	"myapp/helpers"
	"myapp/models"
	"net/http"
	"strconv"
)

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

	return c.JSON(http.StatusOK, helpers.Res("post was deleted"))
}