package main

import (
	"fmt"
	"net/http"
	"url-shortener/const"
	"url-shortener/router"
)

func main() {
	fmt.Println("starting server on port", constants.PORT)
	srv := &http.Server{
		Addr:    constants.PORT,
		Handler: router.Routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
