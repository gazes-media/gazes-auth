package routes

import (
	"encoding/json"
	"gazes-auth/auth"
	"gazes-auth/database"
	"net/http"
)

// ErrorResponse is a struct that is used to send error messages to the client
type ErrorResponse struct {
	Error string `json:"error"`
}

// ErrorResponse.New is a helper function to create a new ErrorResponse
func (e *ErrorResponse) New() *ErrorResponse {
	return &ErrorResponse{}
}

// ErrorResponse.Write is a helper function to write an error response to the client
func (e *ErrorResponse) Write(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e)
}

// Just a little helper function to check if the Authorization header is present
func CheckHeader(w http.ResponseWriter, r *http.Request) bool {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		errJson := ErrorResponse{Error: "No token provided"}
		errJson.Write(w, http.StatusUnauthorized)
		return false
	}
	return true
}

// This function is used to check if the user is authenticated, and if so, return the user.
func AuthGuard(w http.ResponseWriter, r *http.Request) (*database.User, bool) {
	tokenString := r.Header.Get("Authorization")
	if CheckHeader(w, r) == false {
		return nil, false
	}
	// We now need to extract the token from the string because it's a Bearer token
	tokenString = tokenString[7:]
	// validate token
	claims, err := auth.Verify(tokenString)
	if err != nil {
		errJson := ErrorResponse{Error: "Invalid token"}
		errJson.Write(w, http.StatusUnauthorized)
		return nil, false
	}
	// get user from database
	// send user in response
	user := database.User{Email: claims.Email}
	err = user.GetByEmail()
	if err != nil {
		errJson := ErrorResponse{Error: "Error getting user"}
		errJson.Write(w, http.StatusInternalServerError)
		return nil, false
	}
	return &user, true
}
