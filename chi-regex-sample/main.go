package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/{.v[0-9a-zA-Z]{1,}?}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hoge!"))
	})
	http.ListenAndServe(":3000", r)
}
