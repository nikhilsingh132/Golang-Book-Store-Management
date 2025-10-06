package models

import (
	"book_store/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init (){
   config.ConnectToDB()
   db=config.GetDB()
   db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) (*Book, *gorm.DB){
	var book Book
	db:=db.Where("id=?",id).Find(&book)
	return &book, db
}

func DeleteBook(id int64) Book{
	var book Book
	db.Where("id=?",id).Delete(book)
	return book
}

func (b *Book) UpdateBook() *Book {
	db.Save(b)
	return b
}