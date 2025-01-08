package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (cfg *ApiConfig) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// log.Println(string(r.Body))

	// data, _ := io.ReadAll(r.Body)
	// log.Println(string(data))

	type User struct {
		ID string
	}

	var userData User

	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		log.Println("Error in decoding user in delete user function: ", err)
		return
	}
	userIdInt, err := strconv.Atoi(userData.ID)
	if err != nil {
		log.Println("Error in converting user id to int: ", err)
		return
	}

	err = cfg.DBQueries.DeleteUser(context.Background(), int32(userIdInt))
	if err != nil {
		log.Println("Error deleting user from database!! ", err)
		return
	}

	w.Write([]byte(`
	
					<div class="alert alert-danger" role="alert">
تم حذف المستخدم بنجاح

</div>
	`))
}
