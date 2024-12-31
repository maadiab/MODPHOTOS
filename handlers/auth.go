package handlers

import (
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Welcome to MOD")

	w.Header().Set("HX-Redirect", "dashboard")
	w.WriteHeader(http.StatusFound)
	// http.Redirect(w, r, "/dashbaord", http.StatusFound)

}
