package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to create hash : %s", err.Error())
	}
	return string(hash), err
}

func CheckPassword(pass string, hashPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass))
}
