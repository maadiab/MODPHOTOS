package handlers

import (
	"log"
	"net/http"
)

func RegisterSystemUser(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		log.Println(err)
		return
	}

	// name := r.FormValue("name")
	// username := r.FormValue("username")
	// email := r.FormValue("email")
	password := r.FormValue("password")
	reTypePassword := r.FormValue("reTypePassword")

	if password != reTypePassword {
		log.Println("كلمة المرور غير متطابقة")
		return
	}

	// check if there is a user with these information

	// hashing password

	// enter data in database

	// return successfull message in both header and html
}
