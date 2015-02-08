package middleware

import (
  "log"
  "net/http"
  "time"
)

type Middleware struct {
  Next http.Handler
}

func NewWithHandler(handler http.Handler) Middleware {
  return Middleware{handler}
}

func NewWithHandlerFunc(fn http.HandlerFunc) Middleware {
  return NewWithHandler(http.HandlerFunc(fn))
}

func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  t1 := time.Now()
  m.Next.ServeHTTP(w, r)
  t2 := time.Now()
  log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
}
