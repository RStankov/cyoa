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
  RootPath string
  LayoutPath string
}

const staticPath string = "static"
const templatesPath string = "templates"

func New(rootPath string) http.Handler {
  if !strings.HasSuffix(rootPath, "/") {
    rootPath += "/"
  }

  fs := http.FileServer(http.Dir(staticPath))
  handler := http.StripPrefix(rootPath + staticPath + "/", fs)
  layoutPath := path.Join(templatesPath, "layout.html")

  return FileServer{handler, rootPath, layoutPath}
}

func (s FileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if strings.HasPrefix(r.URL.Path, s.RootPath + "/" + staticPath) {
    s.FileHandler.ServeHTTP(w, r)
  } else {
    s.serveTemplate(w, r)
  }
}

func (s FileServer) templatePath(url string) string {
  url = strings.Replace(url, s.RootPath, "", 1)

  if url == "" {
    url = "index"
  }

  return path.Join(templatesPath, url + ".html")
}

func (s FileServer) serveTemplate(w http.ResponseWriter, r *http.Request) {
  fp := s.templatePath(r.URL.Path)

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

  tmpl, err := template.ParseFiles(s.LayoutPath, fp)
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
