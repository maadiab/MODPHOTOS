package handlers

import (
	"net/http"
	"text/template"
)

func ServeRegisterForm(w http.ResponseWriter, r *http.Request) {

	tmp, _ := template.ParseFiles("./templates/register.html")
	tmp.Execute(w, nil)
}

func ServeAddPhotographer(w http.ResponseWriter, r *http.Request) {

	tmp, _ := template.ParseFiles("./templates/addphotographer.html")
	tmp.Execute(w, nil)
}
