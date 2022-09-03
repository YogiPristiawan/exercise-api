package hash

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (hashed string, err error) {
	h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	hashed = string(h)
	return
}

func CompareHashAndPassword(hashPassword string, password string) (err error) {
	rawError := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if rawError != nil {
		err = fmt.Errorf("password salah")
	}
	return
}
