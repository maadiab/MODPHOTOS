package main

import (
	"database/sql"
	"log"
	"net/http"

	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	"github.com/maadiab/modarc/handlers"
	"github.com/maadiab/modarc/internal/database"
)

func main() {

	godotenv.Load()
	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Println("Error in main package: ", err)
	}
	mux := http.NewServeMux()

	server := &http.Server{
		Handler: mux,
		Addr:    "0.0.0.0:8080",
	}

	defer db.Close()
	cfg := &handlers.ApiConfig{
		DBQueries: *database.New(db),
	}

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	mux.Handle("/", http.FileServer(http.Dir("./templates/login")))

	mux.Handle("/dashboard/", http.StripPrefix("/dashboard/", http.FileServer(http.Dir("./templates/dashboard"))))
	// serving forms
	mux.HandleFunc("/reguser", handlers.ServeRegisterForm)
	mux.HandleFunc("/serveusers", handlers.ServeUsersSection)

	mux.HandleFunc("/updatesysuser", handlers.ServeUpdateUserForm)
	mux.HandleFunc("/servephotographers", handlers.ServePhotographersSection)
	mux.HandleFunc("/regphg", handlers.ServeAddPhotographer)
	mux.HandleFunc("/missions", handlers.ServeMissionsSection)
	mux.HandleFunc("/regmission", handlers.ServeAddMission)
	mux.HandleFunc("/overview", handlers.ServeOverView)

	// handlers
	mux.HandleFunc("/regsysuser", cfg.RegisterSystemUser)
	mux.HandleFunc("/users", cfg.GetSysUsers)
	mux.HandleFunc("/updateusers", cfg.UpdateUser)
	mux.HandleFunc("/deleteuser", cfg.DeleteUser)
	mux.HandleFunc("/regphotographer", cfg.AddPhotographer)
	mux.HandleFunc("/photographers", cfg.GetPhotographers)
	mux.HandleFunc("/getmissions", cfg.GetMissons)

	mux.HandleFunc("/api/login", cfg.Login)
	log.Println("Server runnint on: ", server.Addr)
	server.ListenAndServe()
}
