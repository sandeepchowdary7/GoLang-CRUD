package main

import (
	"TestGoProject/src/Mappings"
	"TestGoProject/src/Models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "net/http"
)

var db *gorm.DB

func main() {
	initDb()
	Mappings.InitializeRoutes(db)
}

func initDb() {
	var err error
	db, err = gorm.Open("mysql", "root:helloworld@/testapp2?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Models.Language{}, &Models.BookModel{})
}
