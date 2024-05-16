package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hashedPasssword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPasssword), []byte(password))
	if err != nil {
		return true
	} else {
		return false
	}
}
