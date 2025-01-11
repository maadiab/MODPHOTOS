package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/maadiab/modarc/internal/database"
)

func (cfg *ApiConfig) AddPhotographer(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Error in parsing add photographer form: ", err)
		return
	}

	name := r.FormValue("name")
	username := r.FormValue("username")
	mobile := r.FormValue("mobile")
	password := r.FormValue("password")

	hashedPass, err := HashingPasswords(password)
	if err != nil {
		log.Println("Error in hashing password!!! ", err)
		return
	}

	phg := database.AddPhotographerParams{
		Phgname:     name,
		Phgusername: username,
		Phgmobile:   mobile,
		Phgpassword: hashedPass,
	}

	err = cfg.DBQueries.AddPhotographer(context.Background(), phg)
	if err != nil {
		log.Println("Error at adding photographer to database: ", err)
		return
	}

	w.Write([]byte(`
				<div class="alert alert-success" role="alert">
				تم إضافة المصور بنجاح
	</div>
			`))

}
