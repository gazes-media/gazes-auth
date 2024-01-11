package utils

import (
	"fmt"
	"os"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) string {
	return fmt.Sprintf("%x", argon2.IDKey([]byte(password), []byte(os.Getenv("PASSWORD_SALT")), 1, 64*1024, 4, 32))
}

func VerifyPassword(password, hashedPassword string) bool {
	return HashPassword(password) == hashedPassword
}
