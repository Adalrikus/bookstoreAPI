package controllers

import (
  "fmt"
  "strconv"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/adalrikus/bookstoreAPI/pkg/models"
  utils "github.com/adalrikus/bookstoreAPI/pkg/utils"
)

var newBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
  newBook := models.Book{}
  utils.ParseBody(r, &newBook)
  book := newBook.CreateBook()
  response, _ := json.Marshal(book)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(response)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
  newBooks := models.GetAllBooks()
  response, _ := json.Marshal(newBooks)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(response)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  bookId := vars["bookId"]
  id, err := strconv.ParseInt(bookId, 0, 0)
  if err != nil {
    fmt.Printf("Error getting book: %v\n", err)
  }
  bookDetails, _ := models.GetBookById(id)
  response, _ := json.Marshal(bookDetails)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(response)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
  newBook := models.Book{}
  utils.ParseBody(r, &newBook)
  vars := mux.Vars(r)
  bookId := vars["bookId"]
  id, err := strconv.ParseInt(bookId, 0, 0)
  if err != nil {
    fmt.Printf("Error parsing string: %v\n", err)
  }
  bookDetails, db := models.GetBookById(id)
  if newBook.Title != "" {
    bookDetails.Title = newBook.Title
  }
  if newBook.Author != "" {
    bookDetails.Author = newBook.Author
  }
  if newBook.Publisher != "" {
    bookDetails.Publisher = newBook.Publisher
  }
  db.Save(&bookDetails)
  response, _ := json.Marshal(bookDetails)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(response)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  bookId := vars["bookId"]
  id, err := strconv.ParseInt(bookId, 0, 0)
  if err != nil {
    fmt.Printf("Error getting book: %v\n", err)
  }
  bookDetails := models.DeleteBook(id)
  response, _ := json.Marshal(bookDetails)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(response)
}

