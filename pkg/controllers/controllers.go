package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/adalrikus/bookstoreAPI/pkg/config"
	"github.com/adalrikus/bookstoreAPI/pkg/models"
	utils "github.com/adalrikus/bookstoreAPI/pkg/utils"
)

var queries *models.Queries

func Vars(r *http.Request) map[string]string {
	if rv := r.Context().Value(0); rv != nil {
		return rv.(map[string]string)
	}
	return nil
}

func ResponseWriter(w http.ResponseWriter, r *http.Request, status int, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func LogError(w http.ResponseWriter, r *http.Request, err error) {
	response, _ := json.Marshal(err)
	ResponseWriter(w, r, http.StatusInternalServerError, response)
	log.Fatal(err)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	queries, err := config.Connect(r.Context())
  if err != nil {
    LogError(w, r, err)
  }
	newBook := models.CreateBookParams{}
	if err := utils.ParseBody(r, &newBook); err != nil {
		LogError(w, r, err)
	}
  log.Println(newBook)
	book, err := queries.CreateBook(r.Context(), newBook)
	if err != nil {
		LogError(w, r, err)
	}
	response, err := json.Marshal(book)
	if err != nil {
		log.Fatal(err)
	}
	ResponseWriter(w, r, http.StatusOK, response)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	queries, err := config.Connect(r.Context())
  if err != nil {
    LogError(w, r, err)
  }
	newBooks, err := queries.GetBooks(r.Context())
	if err != nil {
		LogError(w, r, err)
	}
	response, err := json.Marshal(newBooks)
	if err != nil {
		log.Fatal(err)
	}
	ResponseWriter(w, r, http.StatusOK, response)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	queries, err := config.Connect(r.Context())
  if err != nil {
    LogError(w, r, err)
  }
	vars := Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		LogError(w, r, err)
	}
	bookDetails, err := queries.GetBook(r.Context(), id)
	if err != nil {
		LogError(w, r, err)
	}
	response, err := json.Marshal(bookDetails)
	if err != nil {
		log.Fatal(err)
	}
	ResponseWriter(w, r, http.StatusOK, response)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	queries, err := config.Connect(r.Context())
  if err != nil {
    LogError(w, r, err)
  }
	newBook := models.UpdateBookParams{}
	if err := utils.ParseBody(r, &newBook); err != nil {
		LogError(w, r, err)
	}
	if err := queries.UpdateBook(r.Context(), newBook); err != nil {
		LogError(w, r, err)
	}
	response, err := json.Marshal(newBook)
	if err != nil {
		log.Fatal(err)
	}
	ResponseWriter(w, r, http.StatusOK, response)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	queries, err := config.Connect(r.Context())
  if err != nil {
    LogError(w, r, err)
  }
	vars := Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		LogError(w, r, err)
	}
	if err := queries.DeleteBook(r.Context(), id); err != nil {
		LogError(w, r, err)
	}
	response, err := json.Marshal(id)
	if err != nil {
		log.Fatal(err)
	}
	ResponseWriter(w, r, http.StatusOK, response)
}
