package models

import (
	"github.com/SemaMolchanov/go-bookstore/packages/config"
	"gorm.io/gorm"
)

var database *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	database = config.GetDB()
	database.AutoMigrate(&Book{})
}

func (book *Book) CreateBook() *Book {
	database.Create(&book)
	return book
}

func GetAllBooks() []Book {
	var books []Book
	database.Find(&books)
	return books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	database := database.Where("ID=?", Id).Find(&book)
	return &book, database
}

func DeleteBook(ID int64) Book {
	var book Book
	database.Where("ID=?", ID).Delete(book)
	return book
}
