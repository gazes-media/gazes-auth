package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/argon2"
)

// RespondJSON writes the given data as a JSON response to the provided http.ResponseWriter.
// It sets the "Content-Type" header to "application/json" and the HTTP status code to 200 (OK).
// The data parameter should be a struct or a map that can be encoded as JSON.
func RespondJSON(w http.ResponseWriter, data interface{}, statusCode ...int) {
	w.Header().Set("Content-Type", "application/json")
	if len(statusCode) > 0 {
		w.WriteHeader(statusCode[0])
	} else {
		w.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(w).Encode(data)
}

// verify that the environment variables passed in exists
func ValidateEnvVars(envVars []string) {
	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			log.Panicf("Environment variable %s not set", envVar)
		}
	}
}

func HashPassword(password string) string {
	return fmt.Sprintf("%x", argon2.IDKey([]byte(password), []byte(os.Getenv("PASSWORD_SALT")), 1, 64*1024, 4, 32))
}
