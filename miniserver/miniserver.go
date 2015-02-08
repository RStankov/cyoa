package miniserver

import (
  "html/template"
  "log"
  "net/http"
  "os"
  "path"
  "strings"
)

type FileServer struct {
  FileHandler http.Handler
}

func New() http.Handler {
  fs := http.FileServer(http.Dir("static"))
  handler := http.StripPrefix("/static/", fs)

  return FileServer{handler}
}

func (s FileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if strings.HasPrefix(r.URL.Path, "static") {
    s.ServeHTTP(w, r)
  } else {
    serveTemplate(w, r)
  }
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
