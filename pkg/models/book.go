package models

import (
	"github.com/CAndresFernandez/go-books/pkg/config"
	"github.com/jinzhu/gorm"
)

// database points to the ORM
var db *gorm.DB

// the model for a book object
type Book struct {
	gorm.Model
	Name string `gorm:"" json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

// init for the database connection
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// creates a new book
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// findAll books
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// find book by id
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

// delete book by id
func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
