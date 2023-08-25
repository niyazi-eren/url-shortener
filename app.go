package main

import (
	"fmt"
	"net/http"
	"os"
	"url-shortener/router"
)

func main() {
	port := ":80"

	if portEnv := os.Getenv("PORT"); portEnv != "" {
		port = portEnv
	}

	fmt.Println("starting server on port", port)
	srv := &http.Server{
		Addr:    port,
		Handler: router.Routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
