package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashingPassword(p []byte) string {
	password, err := bcrypt.GenerateFromPassword(p, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(password)
}

func ComparePassword(hashedPassword, psw []byte) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(psw))
}