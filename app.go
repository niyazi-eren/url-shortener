package main

import (
	"fmt"
	"net/http"
	"os"
	"url-shortener/router"
)

func main() {
	port := ":8080"

	if portEnv := os.Getenv("PORT"); portEnv != "" {
		port = portEnv
	}

	fmt.Println("starting server on port", port)
	srv := &http.Server{
		Addr:    port,
		Handler: router.Routes(),
	}
	_ = srv.ListenAndServe()
}
