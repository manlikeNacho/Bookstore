package models

import (
	"github.com/jinzhu/gorm"
	"week-1/src/setup"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name     string `gorm:""json:"name"`
	Author   string `json:"author"`
	Language string `json:"language"`
}

func init() {
	setup.Connection()
	db = setup.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) createBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func getAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}
