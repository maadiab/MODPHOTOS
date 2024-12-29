package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	server := &http.Server{
		Handler: mux,
		Addr:    "localhost:8080",
	}

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	mux.Handle("/", http.FileServer(http.Dir("./templates/login")))
	log.Println("Server runnint on: ", server.Addr)
	server.ListenAndServe()
}
