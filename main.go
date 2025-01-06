package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"os"

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
		Addr:    "localhost:8080",
	}

	defer db.Close()
	cfg := &handlers.ApiConfig{
		DBQueries: *database.New(db),
	}

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	mux.Handle("/", http.FileServer(http.Dir("./templates/login")))

	mux.Handle("/dashboard/", http.StripPrefix("/dashboard/", http.FileServer(http.Dir("./templates/dashboard"))))

	mux.HandleFunc("/reguser", handlers.ServeRegisterForm)
	mux.HandleFunc("/serveusers", handlers.ServeUsersSection)
	mux.HandleFunc("/servephotographers", handlers.ServePhotographersSection)
	mux.HandleFunc("/regphotographer", handlers.ServeAddPhotographer)
	mux.HandleFunc("/regmission", handlers.ServeMission)
	mux.HandleFunc("/regsysuser", cfg.RegisterSystemUser)
	mux.HandleFunc("/users", cfg.GetSysUsers)
	mux.HandleFunc("/overview", handlers.ServeOverView)
	mux.HandleFunc("/api/login", cfg.Login)
	log.Println("Server runnint on: ", server.Addr)
	server.ListenAndServe()
}
