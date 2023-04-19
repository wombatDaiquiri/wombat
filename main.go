package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/caddyserver/certmagic"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	metricsRouter := chi.NewRouter()
	metricsRouter.Use(middleware.Timeout(time.Second * 60))
	metricsRouter.Get("/metrics", promhttp.Handler().ServeHTTP)

	metricsServer := &http.Server{
		Addr:         ":42069",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      metricsRouter,
	}
	go func() {
		fmt.Println("run metrics server on 42069")
		metricsErr := metricsServer.ListenAndServe()
		if metricsErr != nil {
			fmt.Println("listen and serve:", metricsErr)
		}
	}()

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(time.Second * 60))

	router.Get("/", serveStatic("content/index.html"))
	router.Get("/favicon.ico", serveStatic("content/static/favicon.ico"))

	router.Get("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./content/static"))).ServeHTTP)

	router.Get("/fonts", serveStatic("content/fonts.html"))

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
		fmt.Println("running with certmagic (http to https redirection port :8080, https port :8081")
		err = certmagic.HTTPS([]string{"daiquiri.fyi", "www.daiquiri.fyi"}, router)
	}

	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("stopping the server")
}
