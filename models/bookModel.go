package models

import (
	"github.com/anujmritunjay/book-management-system/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model

	Name        string `json:"name" gorm:"not null" binding:"required"`
	Author      string `json:"author" gorm:"not null" binding:"required"`
	Publication string `json:"publication" gorm:"not null" binding:"required"`
}

func init() {
	config.ConnectDatabase()
	db = config.DB
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(b)
	return b
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func DeleteBookById(Id int64) Book {
	var book Book
	// First find the book
	db.First(&book, Id) // Changed this line
	// Then delete it
	db.Delete(&book) // Changed this line
	return book
}
