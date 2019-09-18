package service

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Verify(plain string, hashed string) bool {
	byteHash := []byte(hashed)
	bytePlain := []byte(plain)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}