package handlers

import (
	"context"
	"log"
	"net/http"
	// "time"

	"golang.org/x/crypto/bcrypt"
	// "github.com/maadiab/modarc/internal/database"
	// "golang.org/x/crypto/bcrypt"
)

func (cfg *ApiConfig) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Welcome to MOD")

	err := r.ParseForm()

	if err != nil {
		log.Println("Error Parsing Form")
	}

	formUsername := r.FormValue("username")
	formPassword := r.FormValue("password")

	// retrive user from database

	// get user password to compare it with entered
	reqUser, err := cfg.DBQueries.GetUserAuth(context.Background(), formUsername)
	if err != nil {
		log.Println("Error in getting reqUser: ", err)

		return
	}
	// compare it with user data entered

	// b = CompareHash(formPassword, reqUser.Password)
	//
	err = bcrypt.CompareHashAndPassword([]byte(reqUser.Password), []byte(formPassword))
	if err != nil {
		log.Println(err)

		log.Println("Error logging MOD PHOTO SYSTEM!!!")
		w.Write([]byte(`
	
					<div class="alert alert-danger" role="alert">
تم حذف المستخدم بنجاح

</div>
	`))
		return
	} else {

		log.Println("welcome, ", formUsername)
		// time.Sleep(2)
		w.Header().Set("HX-Redirect", "dashboard")
		w.WriteHeader(http.StatusFound)
		// 		w.Write([]byte(`
		// 	<div class="alert alert-success" role="alert">
		// 		User logged in successfully.
		// 	</div>
		// 	<script>
		// 		setTimeout(() => {
		// 			window.location.href = "/dashboard";
		// 		}, 2000); // Redirect after 2 seconds
		// 	</script>
		// `))

		// http.Redirect(w, r, "/dashbaord", http.StatusFound)
		//

	}

	// CompareHash(password, hashedPass)

}
