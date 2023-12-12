package auth

import (
	"fmt"
	"os"

	"golang.org/x/crypto/argon2"
)

// Utility functions for password hashing and comparison
func EncodePassword(password string) string {
	return fmt.Sprintf("%x", argon2.IDKey([]byte(password), []byte(os.Getenv("PASSWORD_SALT")), 1, 64*1024, 4, 32))
}

// ComparePassword compares a password with a hash
func ComparePassword(password string, hash string) bool {
	return EncodePassword(password) == hash
}
