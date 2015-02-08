package main

import (
  "html/template"
  "log"
  "net/http"
  "os"
  "path"
  "encoding/json"
  "github.com/rstankov/cyoa/middleware"
)


type Book struct {
  Id          string  `json:"id"`
  Title       string  `json:"title"`
  Description string  `json:"description"`
  Color       string  `json:"color"`
}

func main() {
  fs := http.FileServer(http.Dir("static"))

  http.Handle("/static/", middleware.NewWithHandler(http.StripPrefix("/static/", fs)))
  http.Handle("/", middleware.NewWithHandlerFunc(serveTemplate))
  http.Handle("/api/", middleware.NewWithHandlerFunc(serveApi))

  log.Println("Listening on 8080...")
  http.ListenAndServe(":8080", nil)
}

func serveApi(w http.ResponseWriter, r *http.Request) {
  book := Book{"1", "Dark river", "The story of ...", "blue"}
  books := []Book{book}

  json.NewEncoder(w).Encode(books)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
  filePath := r.URL.Path

  if filePath == "/" {
    filePath = "index"
  }
  filePath = filePath + ".html"

  lp := path.Join("templates", "layout.html")
  fp := path.Join("templates", filePath)

  // Return a 404 if the template doesn't exist
  info, err := os.Stat(fp)
  if err != nil {
    if os.IsNotExist(err) {
      http.NotFound(w, r)
      return
    }
  }

  // Return a 404 if the request is for a directory
  if info.IsDir() {
    http.NotFound(w, r)
    return
  }

  tmpl, err := template.ParseFiles(lp, fp)
  if err != nil {
    // Log the detailed error
    log.Println(err.Error())
    // Return a generic "Internal Server Error" message
    http.Error(w, http.StatusText(500), 500)
    return
  }

  if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
    log.Println(err.Error())
    http.Error(w, http.StatusText(500), 500)
  }
}
