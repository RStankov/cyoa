package api

import(
  "net/http"
  "encoding/json"
)

type Book struct {
  Id          string  `json:"id"`
  Title       string  `json:"title"`
  Description string  `json:"description"`
  Color       string  `json:"color"`
}

type Api struct {
  RootPath string
}

func New(rootPath string) http.Handler {
  return Api{rootPath};
}

func (s Api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  book := Book{"1", "Dark river", "The story of ...", "blue"}
  books := []Book{book}

  json.NewEncoder(w).Encode(books)
}
