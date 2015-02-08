package main

import (
  "log"
  "net/http"
  "github.com/rstankov/cyoa/api"
  "github.com/rstankov/cyoa/middleware"
  "github.com/rstankov/cyoa/miniserver"
)

func main() {
  http.Handle("/", middleware.NewWithHandler(miniserver.New("/")))
  http.Handle("/api/", middleware.NewWithHandler(api.New("/api/")))

  log.Println("Listening on 8080...")
  http.ListenAndServe(":8080", nil)
}
