package Models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type (
	Language struct {
		gorm.Model
		Name string
	}
	
	FormattedLanguage struct {
		Id uint `json:"id"`
		Name string `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}

	BookModel struct {
		gorm.Model
		Title  string `json:"title"`
		Author string `json:"author"`
		LanguageID int `json:"language_id"`
	}

	FormattedBookModel struct {
		Id        uint       `json:"id"`
		Title     string     `json:"title"`
		Author    string     `json:"author"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}
)
