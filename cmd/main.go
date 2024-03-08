package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
  _ "gorm.io/driver/mysql"
  "github.com/adalrikus/bookstoreAPI/pkg/routes"
)

func main() {
  router := mux.NewRouter()
  routes.RegisterBookStoreRoutes(router)
  http.Handle("/", router)
  log.Fatal(http.ListenAndServe(":8080", router))
}
