package handlers

import (
	"encoding/json"
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

// func ServeUsersList(w http.ResponseWriter, r *http.Request) {
// 	tmp, _ := template.ParseFiles("./templates/userslist.html")
// 	tmp.Execute(w, nil)
// }

func ServePhotographersSection(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("./templates/photographers.html")
	tmp.Execute(w, nil)
}

// serving users forms
func ServeRegisterForm(w http.ResponseWriter, r *http.Request) {

	tmp, _ := template.ParseFiles("./templates/register.html")
	tmp.Execute(w, nil)
}

func ServeUpdateUserForm(w http.ResponseWriter, r *http.Request) {
	// var updateUser database.User

	type userData struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Mobile   string `json:"mobile"`
	}

	var user userData

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// log.Println(r.Body)
		log.Println("Error in decoding update system user form: ", err)
		return
	}
	log.Println("getting data successfully...")
	tmp, err := template.ParseFiles("./templates/edituser.html")
	if err != nil {
		log.Println("Error in parsing edit user template: ", err)
		return
	}
	log.Println(user)
	tmp.Execute(w, user)
}

func ServeAddPhotographer(w http.ResponseWriter, r *http.Request) {

	tmp, _ := template.ParseFiles("./templates/addphotographer.html")
	tmp.Execute(w, nil)
}

// Serve add photographer form
func ServeAddphotographer(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("./template/addphotographer.html")
	tmp.Execute(w, nil)
}

func ServeMission(w http.ResponseWriter, r *http.Request) {

	tmp, _ := template.ParseFiles("./templates/addmission.html")
	tmp.Execute(w, nil)
}
