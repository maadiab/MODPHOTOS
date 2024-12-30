package main

import (
	"log"
	"net/http"

	"github.com/maadiab/modarc/handlers"
)

func main() {

	mux := http.NewServeMux()

	server := &http.Server{
		Handler: mux,
		Addr:    "localhost:8080",
	}

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	mux.Handle("/", http.FileServer(http.Dir("./templates/login")))

	mux.Handle("/dashboard/", http.StripPrefix("/dashboard/", http.FileServer(http.Dir("./templates/dashboard"))))

	mux.HandleFunc("/reguser", handlers.ServeRegisterForm)
	mux.HandleFunc("/regphoto", handlers.ServeAddPhotographer)
	mux.HandleFunc("/api/login", handlers.Login)
	log.Println("Server runnint on: ", server.Addr)
	server.ListenAndServe()
}
