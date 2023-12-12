package auth

import (
	"crypto/sha256"
	"fmt"
)
// Utility functions for password hashing and comparison
func EncodePassword(password string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
}
// ComparePassword compares a password with a hash
func ComparePassword(password string, hash string) bool {
	return EncodePassword(password) == hash
}
