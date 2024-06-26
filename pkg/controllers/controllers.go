package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Adalrikus/bookstoreAPI/pkg/models"
	utils "github.com/Adalrikus/bookstoreAPI/pkg/utils"
)

func ResponseWriter(w http.ResponseWriter, r *http.Request, status int, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(response); err != nil {
		log.Fatal("[ERROR] ", err)
	}
}

func LogError(w http.ResponseWriter, r *http.Request, err error) {
	response, _ := json.Marshal(err)
	ResponseWriter(w, r, http.StatusInternalServerError, response)
	log.Fatal("[ERROR] ", err)
}

func CreateBook(queries *models.Queries) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newBook := models.CreateBookParams{}
		if err := utils.ParseBody(r, &newBook); err != nil {
			LogError(w, r, err)
		}

		book, err := queries.CreateBook(r.Context(), newBook)
		if err != nil {
			LogError(w, r, err)
		}
		log.Println("[INFO] Created book")

		response, err := json.Marshal(book)
		if err != nil {
			log.Fatal("[ERROR] ", err)
		}
		ResponseWriter(w, r, http.StatusOK, response)
	})
}

func GetBooks(queries *models.Queries) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newBooks, err := queries.GetBooks(r.Context())
		if err != nil {
			LogError(w, r, err)
		}
		log.Println("[INFO] Retrieved books")

		response, err := json.Marshal(newBooks)
		if err != nil {
			log.Fatal("[ERROR] ", err)
		}
		ResponseWriter(w, r, http.StatusOK, response)
	})
}

func GetBook(queries *models.Queries) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bookId := r.PathValue("bookId")
		id, err := strconv.ParseInt(bookId, 0, 0)
		if err != nil {
			LogError(w, r, err)
		}

		bookDetails, err := queries.GetBook(r.Context(), id)
		if err != nil {
			LogError(w, r, err)
		}
		log.Println("[INFO] Retrieved book")

		response, err := json.Marshal(bookDetails)
		if err != nil {
			log.Fatal("[ERROR] ", err)
		}
		ResponseWriter(w, r, http.StatusOK, response)
	})
}

func UpdateBook(queries *models.Queries) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		newBook := models.UpdateBookParams{}
		if err := utils.ParseBody(r, &newBook); err != nil {
			LogError(w, r, err)
		}
		bookId := r.PathValue("bookId")
		newBook.ID, err = strconv.ParseInt(bookId, 0, 0)
		if err != nil {
			LogError(w, r, err)
		}
		if err := queries.UpdateBook(r.Context(), newBook); err != nil {
			LogError(w, r, err)
		}
		log.Println("[INFO] Updated book")

		response, err := json.Marshal(newBook)
		if err != nil {
			log.Fatal("[ERROR] ", err)
		}
		ResponseWriter(w, r, http.StatusOK, response)
	})
}

func DeleteBook(queries *models.Queries) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bookId := r.PathValue("bookId")
		id, err := strconv.ParseInt(bookId, 0, 0)
		if err != nil {
			LogError(w, r, err)
		}
		if err := queries.DeleteBook(r.Context(), id); err != nil {
			LogError(w, r, err)
		}
		log.Println("[INFO] Deleted book")

		response, err := json.Marshal(id)
		if err != nil {
			log.Fatal("[ERROR] ", err)
		}
		ResponseWriter(w, r, http.StatusOK, response)
	})
}
