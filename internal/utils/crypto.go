// Package utils provides utility functions
package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string, cost int) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
