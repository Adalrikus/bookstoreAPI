package main

import (
	"log"
  "context"
	"net/http"

	"github.com/Adalrikus/bookstoreAPI/pkg/config"
	"github.com/Adalrikus/bookstoreAPI/pkg/routes"
)

func main() {
  queries, err := config.Connect(context.Background())
  if err != nil {
    log.Fatal("[ERROR] ", err)
  }
	router := http.NewServeMux()
	routes.RegisterBookStoreRoutes(router, queries)
	http.Handle("/", router)
	log.Println("[INFO] Starting server on port 8080")
	log.Fatal("[ERROR] ", http.ListenAndServe(":8080", router))
}
