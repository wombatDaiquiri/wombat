package main

import (
	"flag"
	"fmt"
	"github.com/caddyserver/certmagic"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var isDevelopment = flag.Bool("dev", false, "set to true to run http on port 8080, set to false to run http on 80 and https on 443")

func serveStatic(filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("serving file:", filename)
		http.ServeFile(w, r, filename)
	}
}

func main() {
	flag.Parse()
	fmt.Println("isDevelopment?", isDevelopment != nil && *isDevelopment)
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(time.Second * 60))

	router.Get("/", serveStatic("static/index.html"))
	router.Get("/style.css", serveStatic("static/style.css"))
	router.Get("/favicon.ico", serveStatic("static/favicon.ico"))
	router.Get("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))).ServeHTTP)

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      router,
	}

	var err error
	if isDevelopment != nil && *isDevelopment {
		fmt.Println("run normal server on 8080")
		server.Addr = ":8080"
		err = server.ListenAndServe()
	} else {
		certmagic.DefaultACME.Agreed = true
		certmagic.DefaultACME.Email = "wombatDaiquiri@proton.me"
		certmagic.DefaultACME.CA = certmagic.LetsEncryptProductionCA
		certmagic.Default.Storage = &certmagic.FileStorage{Path: "/certs"}

		certmagic.HTTPPort, certmagic.HTTPSPort = 8080, 8081
		fmt.Println("running with certmagic")
		err = certmagic.HTTPS([]string{"daiquiri.fyi", "www.daiquiri.fyi"}, router)
	}

	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("stopping the server")
}
