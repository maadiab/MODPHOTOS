package handlers

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Welcome to MOD")

	err := r.ParseForm()

	if err != nil {
		log.Println("Error Parsing Form")
	}

	username := r.FormValue("username")
	// password := r.FormValue("password")

	// retrive user from database

	// compare it with user data entered
	// CompareHash(password, hashedPass)

	w.Header().Set("HX-Redirect", "dashboard")
	w.WriteHeader(http.StatusFound)
	// http.Redirect(w, r, "/dashbaord", http.StatusFound)
	log.Println("welcome, ", username)
}

func CompareHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(password))
	if err != nil {
		log.Println(err)
	}
	return err == nil
}
