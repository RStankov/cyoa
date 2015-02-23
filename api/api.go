package api

import(
  "net/http"
  "encoding/json"
  "database/sql"
  "log"
  "strconv"
  "regexp"
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

type ApiError struct {
  Code        int     `json:id`
  Description string  `json:description`
}

type Api struct {
  RootPath string
}

func New(rootPath string) http.Handler {
  return Api{rootPath};
}

func (s Api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  bookReg, _ := regexp.Compile("^/api/books/([0-9]+)$")

  if r.Method == "POST" && r.URL.Path == "/api/books" {
    r.ParseForm()

    book := Book{0, r.FormValue("title"), r.FormValue("description"), r.FormValue("color")}

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
  } else if r.Method == "GET" && r.URL.Path == "/api/books" {
    limit := extractLimit(r.URL.Query()["limit"], 50)

    db, err := sql.Open("neo4j-cypher", "http://192.168.59.103:7474")
    if err != nil {
      log.Fatal(err)
    }
    defer db.Close()

    stmt, err := db.Prepare("MATCH (r:Book) RETURN id(r), r.title, r.description, r.color LIMIT {0}")
    if err != nil {
      log.Fatal(err)
    }

    rows, err := stmt.Query(limit)
    if err != nil {
      log.Fatal(err)
    }
    defer rows.Close()

    books := make([]Book, 0)

    for rows.Next() {
      id := 0
      title := ""
      description := ""
      color := ""

      err = rows.Scan(&id, &title, &description, &color)
      if err != nil {
        log.Fatal(err)
      }

      books = append(books, Book{id, title, description, color})
    }

    json.NewEncoder(w).Encode(books)
  } else if r.Method == "GET" && bookReg.MatchString(r.URL.Path) {
    id, _ := strconv.Atoi(bookReg.FindStringSubmatch(r.URL.Path)[1])

    db, err := sql.Open("neo4j-cypher", "http://192.168.59.103:7474")
    if err != nil {
      log.Fatal(err)
    }
    defer db.Close()

    stmt, err := db.Prepare("MATCH (r:Book) WHERE id(r) = {0} RETURN id(r), r.title, r.description, r.color")
    if err != nil {
      log.Fatal(err)
    }

    rows, err := stmt.Query(id)
    if err != nil {
      log.Fatal(err)
    }
    defer rows.Close()

    rows.Next()

    title := ""
    description := ""
    color := ""

    err = rows.Scan(&id, &title, &description, &color)
    if err != nil {
      log.Fatal(err)
    }

    book := Book{id, title, description, color}

    json.NewEncoder(w).Encode(book)
  } else if r.Method == "DELETE" && bookReg.MatchString(r.URL.Path) {
    id, _ := strconv.Atoi(bookReg.FindStringSubmatch(r.URL.Path)[1])

    db, err := sql.Open("neo4j-cypher", "http://192.168.59.103:7474")
    if err != nil {
      log.Fatal(err)
    }
    defer db.Close()

    stmt, err := db.Prepare("MATCH (r:Book) WHERE id(r) = {0} DELETE r")
    if err != nil {
      log.Fatal(err)
    }

    rows, err := stmt.Query(id)
    if err != nil {
      log.Fatal(err)
    }
    defer rows.Close()

    w.WriteHeader(204)
  } else {
    w.WriteHeader(404)
    json.NewEncoder(w).Encode(ApiError{404, "Not Found"})
  }
}

func extractLimit(limitQuery []string, maxLimit int) int {
  if len(limitQuery) == 0 {
    return maxLimit
  }

  limit, err := strconv.Atoi(limitQuery[0])

  if err != nil {
    return maxLimit
  }

  return limit
}
