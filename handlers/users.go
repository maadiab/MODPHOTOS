package handlers

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/maadiab/modarc/internal/database"
)

func (cfg *ApiConfig) RegisterSystemUser(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		log.Println(err)
		return
	}

	name := r.FormValue("name")
	username := r.FormValue("username")
	email := r.FormValue("email")
	mobile := r.FormValue("mobile")
	password := r.FormValue("password")
	reTypePassword := r.FormValue("reTypePassword")

	if password != reTypePassword {
		log.Println("كلمة المرور غير متطابقة")
		return
	}

	// log.Println(mobile)
	// check if there is a user with these information

	// hashing password
	hashedPass, err := HashingPasswords(password)
	if err != nil {
		log.Println("Error with hashing at registering new users: ", err)
		return
	}
	// enter data in database

	nuser := database.CreateUserParams{
		Name:     name,
		Username: username,
		Email:    email,
		Mobile: sql.NullString{
			String: mobile,
			Valid:  mobile != "",
		},
		Password: hashedPass,
	}

	// log.Println(nuser.Mobile.String)
	err = cfg.DBQueries.CreateUser(context.Background(), nuser)
	if err != nil {
		log.Println("Error in RegisterSysUser: ", err)
		return
	}

	w.Write([]byte("it's ok"))
	// return successfull message in both header and html
}
