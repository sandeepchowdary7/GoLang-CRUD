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
	_ = router.Run()
}
