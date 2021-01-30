// github.com/bartmika/mulberry-server/pkg/utils/password.go
package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Function takes the plaintext string and returns a hash string.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Function checks the plaintext string and hash string and returns either true
// or false depending.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
