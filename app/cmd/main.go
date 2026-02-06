package main

import (
	"log"
	"net/http"

	"github.com/jl54/jonasluethi-web/internal/handlers"
)

func main() {
	pageHandler := &handlers.PageHandler{}

	router := http.NewServeMux()
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	router.HandleFunc("/{$}", pageHandler.Handle)

	server := &http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
