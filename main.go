package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "path - %s", r.URL.Path[1:])
    fmt.Printf("GET %s\n", r.URL.Path)
}

func main() {
    fmt.Println("Server running at port 8080")

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
