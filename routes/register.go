package routes

import (
	"encoding/json"
	"gazes-auth/auth"
	"net/http"
	"gazes-auth/database"
)

type RegisterResponse struct {
	Token string `json:"token"`
}


func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	r.Body.Close();
	// unmarshal request body
	var user database.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check if user exists
	u := database.User{Email: user.Email, Username: user.Username, Password: auth.EncodePassword(user.Password)}
	err = u.GetByEmail()
	if err == nil {
		jsonErr := ErrorResponse{Error: "User already exists"}
		jsonErr.Write(w, http.StatusUnauthorized)
		return
	}

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