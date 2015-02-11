package api

import(
  "net/http"
  "encoding/json"
  "database/sql"
  "log"
  _ "gopkg.in/cq.v1"
)

type Book struct {
  Id          int     `json:"id"`
  Title       string  `json:"title"`
  Description string  `json:"description"`
  Color       string  `json:"color"`
}

type Page struct {
  Id        int       `json:"id"`
  Text      string    `json:"text"`
  Choices   []Choice  `json:"choices"`
}

type Choice struct {
  Id         int      `json:"id"`
  Text       string   `json:"text"`
  NextPageId int      `json:"nextPageId"`
}

type Api struct {
  RootPath string
}

func New(rootPath string) http.Handler {
  return Api{rootPath};
}

func (s Api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" && r.URL.String() == "/api/books" {
    book := Book{0, "Dark river", "The story of ...", "blue"}

    db, err := sql.Open("neo4j-cypher", "http://192.168.59.103:7474")
    if err != nil {
      log.Fatal(err)
    }
    defer db.Close()

    stmt, err := db.Prepare("CREATE (record:Book {title:{0}, description: {1}, color: {2}}) RETURN id(record)")
    if err != nil {
      log.Fatal(err)
    }

    rows, err := stmt.Query(book.Title, book.Description, book.Color)
    if err != nil {
      log.Fatal(err)
    }
    defer rows.Close()

    rows.Next()
    err = rows.Scan(&book.Id)
    if err != nil {
      log.Fatal(err)
    }

    json.NewEncoder(w).Encode(book)
  } else {
    json.NewEncoder(w).Encode("Not found")
  }
}
