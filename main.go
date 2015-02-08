package main

import (
  "html/template"
  "log"
  "net/http"
  "os"
  "path"
  "time"
  "encoding/json"
)

type Middleware struct {
  Next http.Handler
}

func withMiddleware(handler http.Handler) Middleware {
  return Middleware{handler}
}

func withMiddlewareFunc(fn http.HandlerFunc) Middleware {
  return withMiddleware(http.HandlerFunc(fn))
}

func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  t1 := time.Now()
  m.Next.ServeHTTP(w, r)
  t2 := time.Now()
  log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
}

type Book struct {
  Id          string  `json:"id"`
  Title       string  `json:"title"`
  Description string  `json:"description"`
  Color       string  `json:"color"`
}

func main() {
  fs := http.FileServer(http.Dir("static"))

  http.Handle("/static/", withMiddleware(http.StripPrefix("/static/", fs)))
  http.Handle("/", withMiddlewareFunc(serveTemplate))
  http.Handle("/api/", withMiddlewareFunc(serveApi))

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
