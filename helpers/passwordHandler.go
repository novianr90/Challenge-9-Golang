package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(param string) string {
	var (
		salt     = 8
		password = []byte(param)
	)

	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hash)
}

func CompareHashAndPass(h, p []byte) bool {
	err := bcrypt.CompareHashAndPassword(h, p)

	return err == nil
}
