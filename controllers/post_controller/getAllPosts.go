package post_controller

import (
	"encoding/xml"
	"github.com/labstack/echo/v4"
	"myapp/constants"
	"myapp/database"
	"myapp/helpers"
	"myapp/models"
	"net/http"
)

func getAllPosts(c echo.Context, responseType string) error {
	var postsFromDB []models.Post
	getPostRes := database.DB.Find(&postsFromDB)
	if getPostRes.Error != nil {
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

// retrieves post from db
func GetAllPostsXML(c echo.Context) error {
	return getAllPosts(c, "XML")
}

// retrieves post from db
func GetAllPostsJSON(c echo.Context) error {
	return getAllPosts(c, "JSON")
}
