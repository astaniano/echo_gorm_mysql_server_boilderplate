package post_controller

import (
	"github.com/labstack/echo/v4"
	"myapp/constants"
	"myapp/helpers"
	"myapp/models"
	"net/http"
	"strconv"
)

func GetPost(c echo.Context) error {
	responseType := c.QueryParam("type")
	if responseType != constants.JSON && responseType != constants.XML {
		return c.JSON(http.StatusBadRequest, helpers.Res("unknown type of response"))
	}

	postIdFromReq, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res("postID was not in the url"))
	}

	var postModel models.Post
	postFromDB, err := postModel.GetPostByID(postIdFromReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res("could not find post for the specified id"))
	}

	if responseType == constants.XML {
		return c.XML(http.StatusOK, postFromDB)
	}
	return c.JSON(http.StatusOK, postFromDB)
}
