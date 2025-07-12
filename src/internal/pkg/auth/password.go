package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidPassword = errors.New("invalid password")
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, has string) error {
	return bcrypt.CompareHashAndPassword([]byte(has), []byte(password))
}