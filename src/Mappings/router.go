package Mappings

import (
	"TestGoProject/src/Controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitializeRoutes(db *gorm.DB) {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	router.GET("/", Controllers.Index)
	router.POST("/createBook", Controllers.Create)
	router.PATCH("/book/:id", Controllers.Update)
	router.DELETE("/book/:id", Controllers.Delete)
	router.GET("/book/:id", Controllers.Fetch)
	router.POST("/createLanguage", Controllers.CreateLanguage)
	router.GET("/languages", Controllers.IndexLanguages)
	router.DELETE("/language/:id", Controllers.DeleteLanguage)
	router.GET("/language/:id", Controllers.FetchLanguage)
	router.PATCH("/language/:id", Controllers.UpdateLanguage)
	_ = router.Run()
}
