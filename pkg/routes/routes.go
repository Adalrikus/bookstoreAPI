package routes

import (
	"net/http"

	"github.com/Adalrikus/bookstoreAPI/pkg/controllers"
	"github.com/Adalrikus/bookstoreAPI/pkg/models"
)

func RegisterBookStoreRoutes(router *http.ServeMux, queries *models.Queries) {
	router.HandleFunc("POST /book/", controllers.CreateBook(queries))
	router.HandleFunc("GET /book/", controllers.GetBooks(queries))
	router.HandleFunc("GET /book/{bookId}", controllers.GetBook(queries))
	router.HandleFunc("PUT /book/{bookId}", controllers.UpdateBook(queries))
	router.HandleFunc("DELETE /book/{bookId}", controllers.DeleteBook(queries))
}
