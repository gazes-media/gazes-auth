package routes

import (
	"encoding/json"
	"gazes-auth/database"
	"net/http"
)

type UserMeResponse struct {
	User database.User `json:"user"`
}

func UserMeHandler(w http.ResponseWriter, r *http.Request) {
	// get Token from request header
	user, ok := AuthGuard(w, r)
	if !ok {
		jsonErr := ErrorResponse{Error: "Error getting user"}
		jsonErr.Write(w, http.StatusInternalServerError)
		return
	}
	userModified := database.User{Email: user.Email, Username: user.Username, Model: user.Model}
	// send user in response, but without the password
	response := UserMeResponse{User: userModified}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
