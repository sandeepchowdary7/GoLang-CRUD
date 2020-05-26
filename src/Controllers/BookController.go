package Controllers

import (
	"TestGoProject/src/Models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Index(c *gin.Context) {
	dbb := c.MustGet("db").(*gorm.DB)
	var books []Models.BookModel
	var _books []Models.FormattedBookModel

	dbb.Find(&books)
	if len(books) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "No data Found"})
		return
	}

	for _, item := range books {
		_books = append(_books, Models.FormattedBookModel{Id: item.ID, Title: item.Title, Author: item.Author, CreatedAt: item.CreatedAt, UpdatedAt: item.UpdatedAt, DeletedAt: item.DeletedAt})
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _books})
}

func Fetch(c *gin.Context)  {
	dbb := c.MustGet("db").(*gorm.DB)
	var book Models.BookModel
	bookId := c.Param("id")

	dbb.Find(&book, bookId)

	if book.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "No Book Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "book": book})
}

func Create(c *gin.Context) {
	dbb := c.MustGet("db").(*gorm.DB)
	book := Models.BookModel{
		Title:  c.PostForm("title"),
		Author: c.PostForm("author"),
	}

	dbb.Save(&book)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Book created successfully!", "bookId": book.ID})
}

func Update(c *gin.Context)  {
	dbb := c.MustGet("db").(*gorm.DB)
	var book Models.BookModel
	bookId := c.Param("id")

	dbb.First(&book, bookId)

	if book.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "Book Not Found"})
		return
	}

	dbb.Model(&book).Updates(map[string]interface{}{"title": c.PostForm("title"), "author": c.PostForm("author")})
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Book Updated Successfully!"})
}

func Delete(c *gin.Context)  {
	dbb := c.MustGet("db").(*gorm.DB)
	var book Models.BookModel
	bookId := c.Param("id")

	dbb.Find(&book, bookId)

	if book.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Book Found"})
		return
	}

	dbb.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"status":http.StatusOK, "message": "Book Deleted Successfully"})
}
