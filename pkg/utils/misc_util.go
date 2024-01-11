package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func Getenvars(keys []string) (map[string]string, error) {
	values := make(map[string]string)
	missing := []string{}

	for i := range keys {
		value := os.Getenv(keys[i])

		if value == "" {
			missing = append(missing, keys[i])
		} else {
			values[keys[i]] = value
		}
	}

	if len(missing) > 0 {
		return values, fmt.Errorf("missing environment variables: %v", missing)
	}

	return values, nil
}

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
