package handlers

import (
	"context"
	"log"
	"net/http"
	"text/template"
	// "github.com/maadiab/modarc/internal/database"
	// "github.com/maadiab/modarc/internal/da;tabase"
)

func (cfg *ApiConfig) GetPhotographers(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("photographers list"))
	// define []struct to hold data
	// var users []database.Photographer

	// bring photographers from database

	phgs, err := cfg.DBQueries.GetAllPHotographers(context.Background())

	if err != nil {
		log.Println("Error getting photograhpers from database", err)
		return
	}
	//
	//
	//
	// parse data into template
	tmp, err := template.ParseFiles("./templates/photographerslist.html")
	if err != nil {
		log.Println("Error in parsing photographers list: ", phgs)
		return
	}
	log.Println(phgs)
	tmp.Execute(w, phgs)
}
