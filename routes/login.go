package routes

import (
	"encoding/json"
	"gazes-auth/auth"
	"net/http"
	"gazes-auth/database"
)


type LoginResponse struct {
	Token string `json:"token"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	r.Body.Close();
	// unmarshal request body
	var user auth.UserLogin
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check if user exists
	u := database.User{Email: user.Email}
	err = u.GetByEmail()
	if err != nil {
		jsonErr := ErrorResponse{Error: "User does not exist"}
		jsonErr.Write(w, http.StatusUnauthorized)
		return
	}

	// check if password is correct
	if !auth.ComparePassword(u.Password, user.Password) {
		jsonErr := ErrorResponse{Error: "Invalid password"}
		jsonErr.Write(w, http.StatusUnauthorized)
		return
	}

	// create JWT token
	token, err := auth.Sign(u)
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