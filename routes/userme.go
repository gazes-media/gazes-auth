package routes

import (
	"encoding/json"
	"gazes-auth/database"
	"net/http"
)

type UserMeResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       uint   `json:"id"`
}

func UserMeHandler(w http.ResponseWriter, r *http.Request, user *database.User) {
	userModified := database.User{Email: user.Email, Username: user.Username, Model: user.Model}
	// send user in response, but without the password
	response := UserMeResponse{Username: userModified.Username, Email: userModified.Email, ID: userModified.ID}
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}
