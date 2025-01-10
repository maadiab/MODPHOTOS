package handlers

import (
	"log"
	"net/http"
	"text/template"
	// "github.com/maadiab/modarc/internal/da;tabase"
)

func (cfg *ApiConfig) GetPhotographers(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("photographers list"))
	// define []struct to hold data
	// var users []database.Photographerc
	// bring users from database
	//
	//
	//
	//
	// parse data into template
	tmp, err := template.ParseFiles("./templates/photographerslist.html")
	if err != nil {
		log.Println("Error in parsing photographers list: ", err)
		return
	}

	tmp.Execute(w, nil)
}
