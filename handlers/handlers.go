package handlers

import (
	"log"

	"github.com/maadiab/modarc/internal/database"
	"golang.org/x/crypto/bcrypt"
)

type ApiConfig struct {
	DBQueries database.Queries
}

func HashingPasswords(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println("Error with hashing password: ", err)
		return "", err
	}

	return string(bytes), err
}
