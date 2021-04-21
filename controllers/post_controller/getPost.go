package post_controller

import (
	"github.com/labstack/echo/v4"
	"myapp/constants"
	"myapp/helpers"
	"myapp/models"
	"net/http"
	"strconv"
)

func getPost(c echo.Context, responseType string) error {
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

// retrieves post from db
func GetPostXML(c echo.Context) error {
	return getPost(c, "XML")
}

// retrieves post from db
func GetPostJSON(c echo.Context) error {
	return getPost(c, "JSON")
}
