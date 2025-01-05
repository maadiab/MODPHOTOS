package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func ServeOverView(w http.ResponseWriter, r *http.Request) {

	tmp, err := template.ParseFiles("./templates/overview.html")
	if err != nil {
		log.Println("Error parsing overview template!!!")
		return
	}
	tmp.Execute(w, nil)
}

func ServeUsersSection(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("./templates/users.html")
	tmp.Execute(w, nil)
}

func ServePhotographersSection(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("./templates/photographers.html")
	tmp.Execute(w, nil)
}

func ServeRegisterForm(w http.ResponseWriter, r *http.Request) {

	tmp, _ := template.ParseFiles("./templates/register.html")
	tmp.Execute(w, nil)
}

func ServeAddPhotographer(w http.ResponseWriter, r *http.Request) {

	tmp, _ := template.ParseFiles("./templates/addphotographer.html")
	tmp.Execute(w, nil)
}

func ServeMission(w http.ResponseWriter, r *http.Request) {

	tmp, _ := template.ParseFiles("./templates/addmission.html")
	tmp.Execute(w, nil)
}
