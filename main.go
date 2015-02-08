package main

import (
  "log"
  "net/http"
  "encoding/json"
  "github.com/rstankov/cyoa/middleware"
  "github.com/rstankov/cyoa/miniserver"
)


type Book struct {
  Id          string  `json:"id"`
  Title       string  `json:"title"`
  Description string  `json:"description"`
  Color       string  `json:"color"`
}

func main() {
  http.Handle("/", middleware.NewWithHandler(miniserver.New("/")))
  http.Handle("/api/", middleware.NewWithHandlerFunc(serveApi))

  log.Println("Listening on 8080...")
  http.ListenAndServe(":8080", nil)
}

func serveApi(w http.ResponseWriter, r *http.Request) {
  book := Book{"1", "Dark river", "The story of ...", "blue"}
  books := []Book{book}

  json.NewEncoder(w).Encode(books)
}
