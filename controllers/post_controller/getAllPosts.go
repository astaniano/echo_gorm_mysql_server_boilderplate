package post_controller

import (
	"encoding/xml"
	"github.com/labstack/echo/v4"
	"myapp/constants"
	"myapp/helpers"
	"myapp/models"
	"net/http"
)

func GetAllPosts(c echo.Context) error {
	responseType := c.QueryParam("type")
	if responseType != constants.JSON && responseType != constants.XML {
		return c.JSON(http.StatusBadRequest, helpers.Res("unknown type of response"))
	}

	var postModel models.Post
	postsFromDB, err := postModel.GetAllPosts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res("could not get posts from db"))
	}

	if responseType == constants.XML {
		type allPosts struct {
			XMLName xml.Name      `xml:"posts"`
			Posts   []models.Post `xml:"post"`
		}
		return c.XML(http.StatusOK, &allPosts{Posts: postsFromDB})
	}
	return c.JSON(http.StatusOK, &postsFromDB)
}
