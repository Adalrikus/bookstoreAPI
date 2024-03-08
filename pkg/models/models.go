package models

import (
  "gorm.io/gorm"
  "github.com/adalrikus/bookstoreAPI/pkg/config"
)

var DB *gorm.DB

type Book struct {
  gorm.Model
  Title string `json:"title"`
  Author string `json:"author"`
  Publisher string `json:"publisher"`
}

func init() {
  config.ConnectDB()
  DB = config.GetDB()
  DB.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
  DB.Create(&b)
  return b
}

func GetAllBooks() []Book {
  var books []Book
  DB.Find(&books)
  return books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
  var book Book
  DB := DB.Where("id = ?", id).Find(&book)
  return &book, DB
}

func UpdateBook(book *Book) Book {
  DB.Save(&book)
  return *book
}

func DeleteBook(id int64) Book {
  var book Book
  DB.Where("ID=?", id).Delete(book)
  return book
}
