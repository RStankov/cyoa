package main

import (
    "log"
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "path - %s", r.URL.Path[1:])
    fmt.Printf("GET %s\n", r.URL.Path)
}

func main() {
    fs := http.FileServer(http.Dir("static"))

    log.Println("Listening on 8080...")

    http.Handle("/", fs)
    http.HandleFunc("/api/", handler)
    http.ListenAndServe(":8080", nil)
}
