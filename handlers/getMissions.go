package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func (cfg *ApiConfig) GetMissons(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("./templates/missionslist.html")
	if err != nil {
		log.Println("Error in parsing missions list!!!, ", err)
		return
	}

	tmp.Execute(w, nil)
}
