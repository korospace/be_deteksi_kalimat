package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(passwordHash)
}

func VerifyPassword(hasPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hasPassword), []byte(password))

	return err
}
