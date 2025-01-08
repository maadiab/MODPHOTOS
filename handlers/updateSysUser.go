package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/maadiab/modarc/internal/database"
)

func (cfg *ApiConfig) UpdateUser(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println("Error in parsing update user form: ", err)
		return
	}

	id := r.FormValue("id")
	name := r.FormValue("name")
	username := r.FormValue("username")
	email := r.FormValue("email")
	mobile := r.FormValue("mobile")

	numId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error in converting id to int: ", err)
		return
	}

	uuser := database.UpdateUserParams{
		ID:       int32(numId),
		Name:     name,
		Username: username,
		Email:    email,
		Mobile:   mobile,
	}
	err = cfg.DBQueries.UpdateUser(context.Background(), uuser)

	if err != nil {
		log.Println("Error in Update User in database function: ", err)
		return
	}

	w.Write([]byte(`
	
				<div class="alert alert-success" role="alert">
تم تحديث بيانات المستخدم بنجاح	</div>
	`))
}
