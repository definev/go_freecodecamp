package models

import (
	"github.com/definev/go_freecodecamp/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name,omitempty"`
	Author      string `json:"author,omitempty"`
	Publication string `json:"publication,omitempty"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) (*gorm.DB, *Book) {
	var getBook Book
	db := db.Where("ID=?", id).Find(&getBook)
	return db, &getBook
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(book)
	return book
}
