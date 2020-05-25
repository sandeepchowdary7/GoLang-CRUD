package Models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type (
	BookModel struct {
		gorm.Model
		Title  string `json:"title"`
		Author string `json:"author"`
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
