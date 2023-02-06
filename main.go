package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func serveStatic(filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	}
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Second * 60))

	r.Get("/", serveStatic("static/index.html"))
	r.Get("/style.css", serveStatic("static/style.css"))
	r.Get("/favicon.ico", serveStatic("static/favicon.ico"))
	r.Get("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))).ServeHTTP)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("stopping the server")
}
