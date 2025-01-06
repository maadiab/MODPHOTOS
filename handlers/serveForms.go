package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"text/template"

	"github.com/maadiab/modarc/internal/database"
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
	var updateUser database.User

	body, _ := io.ReadAll(r.Body)

	err := json.Unmarshal(body, &updateUser)
	if err != nil {
		log.Println("Error in decoding update system user form using unmarshall: ", err)
		return
	}
	// log.Println(string(body))

	err = json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		log.Println(r.Body)
		log.Println("Error in decoding update system user form: ", err)
		return
	}
	tmp, err := template.ParseFiles("./templates/edituser.html")
	if err != nil {
		log.Println("Error in parsing edit user template: ", err)
		return
	}
	tmp.Execute(w, updateUser)
}

func ServeAddPhotographer(w http.ResponseWriter, r *http.Request) {

	tmp, _ := template.ParseFiles("./templates/addphotographer.html")
	tmp.Execute(w, nil)
}

func ServeMission(w http.ResponseWriter, r *http.Request) {

	tmp, _ := template.ParseFiles("./templates/addmission.html")
	tmp.Execute(w, nil)
}
