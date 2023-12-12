package routes

import (
	"encoding/json"
	"gazes-auth/auth"
	"gazes-auth/database"
	"net/http"

	"gopkg.in/validator.v2"
)

type RegisterResponse struct {
	Token string `json:"token"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	// unmarshal request body
	var userRegister auth.UserRegister
	err := json.NewDecoder(r.Body).Decode(&userRegister)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// validate request body fields
	err = validator.Validate(userRegister)
	if err != nil {
		jsonErr := ErrorResponse{Error: "Invalid request body: " + err.Error()}
		jsonErr.Write(w, http.StatusBadRequest)
		return
	}

	// check if user exists
	user := database.User{Email: userRegister.Email, Username: userRegister.Username, Password: auth.EncodePassword(userRegister.Password)}
	// create user
	err = user.Create()
	if err != nil {
		jsonErr := ErrorResponse{Error: "Error creating user"}
		jsonErr.Write(w, http.StatusInternalServerError)
		return
	}

	// create JWT token
	token, err := auth.Sign(user)
	if err != nil {
		jsonErr := ErrorResponse{Error: "Error creating token"}
		jsonErr.Write(w, http.StatusInternalServerError)
		return
	}

	// send token in response
	response := LoginResponse{Token: token}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
