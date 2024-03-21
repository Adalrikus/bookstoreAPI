package main

import (
	"log"
	"net/http"

	"github.com/adalrikus/bookstoreAPI/pkg/routes"
)

func main() {
	router := http.NewServeMux()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	log.Println("[INFO] Starting server on port 8080")
	log.Fatal("[ERROR] ", http.ListenAndServe(":8080", router))
}
