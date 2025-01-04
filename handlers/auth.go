package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/maadiab/modarc/internal/database"
	// "golang.org/x/crypto/bcrypt"
)

func (cfg *ApiConfig) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Welcome to MOD")

	err := r.ParseForm()

	if err != nil {
		log.Println("Error Parsing Form")
	}

	formUsername := r.FormValue("username")
	formPassword := r.FormValue("password")

	// retrive user from database

	user := &database.GetUserAuthParams{
		Username: formUsername,
		Password: formPassword,
	}

	// reqUser := database.User{}

	reqUser, err := cfg.DBQueries.GetUserAuth(context.Background(), *user)

	if err != nil {
		log.Println("Error in getting reqUser: ", err)
		return
	}

	// compare it with user data entered

	CompareHash(formPassword, reqUser.Password)
	// CompareHash(password, hashedPass)

	w.Header().Set("HX-Redirect", "dashboard")
	w.WriteHeader(http.StatusFound)
	// http.Redirect(w, r, "/dashbaord", http.StatusFound)
	log.Println("welcome, ", formUsername, reqUser)
}
