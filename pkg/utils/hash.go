package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", err
	}

	newHashPassword := string(bytes)

	return newHashPassword, nil
}

func CheckPassword(authPassword, userPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(authPassword))
	if err != nil {
		return err
	}
	return nil
}
