package routes

import (
	"net/http"

	"github.com/Adalrikus/bookstoreAPI/pkg/controllers"
)

func RegisterBookStoreRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /book/", controllers.CreateBook)
	router.HandleFunc("GET /book/", controllers.GetBooks)
	router.HandleFunc("GET /book/{bookId}", controllers.GetBook)
	router.HandleFunc("PUT /book/{bookId}", controllers.UpdateBook)
	router.HandleFunc("DELETE /book/{bookId}", controllers.DeleteBook)
}
