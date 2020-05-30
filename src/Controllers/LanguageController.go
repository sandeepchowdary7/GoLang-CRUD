package Controllers

import (
	"TestGoProject/src/Models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func IndexLanguages(c *gin.Context)  {
	dbb := c.MustGet("db").(*gorm.DB)
	var languages []Models.Language
	var _languages []Models.FormattedLanguage

	dbb.Find(&languages)
	if len(languages) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "No Languages Found!"})
		return
	}

	for _, item := range languages {
		_languages = append(_languages, Models.FormattedLanguage{
			Id:       item.ID,
			Name:      item.Name,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
			DeletedAt: item.DeletedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "languages": _languages})

}

func CreateLanguage(c *gin.Context)  {
	dbb := c.MustGet("db").(*gorm.DB)
	language := Models.Language{
		Name: c.PostForm("name"),
	}

	if language.Name != "" {
		dbb.Save(&language)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Language Created.", "language": language.ID})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusNotFound, "data": "Language Name is Empty."})
		return
	}
}