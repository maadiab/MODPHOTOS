package handlers

import (
	"context"
	"log"
	"net/http"
	"text/template"
)

func (cfg *ApiConfig) GetSysUsers(w http.ResponseWriter, r *http.Request) {

	users, err := cfg.DBQueries.GetAllUsers(context.Background())
	if err != nil {
		log.Println("Error in Getting SysUsers: ", err)
	}

	tmp, err := template.ParseFiles("./templates/userslist.html")
	if err != nil {
		log.Println("Error parsing userslist template: ", err)
	}

	tmp.Execute(w, users)

	// for _, user := range users {
	// 	log.Println(user.Name)
	// }

	// log.Println(users)
}
