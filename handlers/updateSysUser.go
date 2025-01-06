package handlers

import (
	// "encoding/json"
	"log"
	"net/http"
	"text/template"
)

func (cfg *ApiConfig) UpdateSysUser(w http.ResponseWriter, r *http.Request) {

	// var user UserData

	// err := json.NewDecoder(r.Body).Decode(&user)
	// if err != nil {
	// 	log.Println("Error in decoding user for update: ", err)
	// 	return
	// }

	tmp, err := template.ParseFiles("./templates/edituser.html")
	if err != nil {
		log.Println("Error in parsing edituser form: ", err)
		return
	}
	tmp.Execute(w, nil)
}
