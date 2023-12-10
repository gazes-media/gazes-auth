package routes

import (
	"encoding/json"
	"gazes-auth/auth"
	"net/http"
)


func LoginHandler(w http.ResponseWriter, r *http.Request) {

	r.Body.Close();
	// unmarshal request body
	var user auth.UserLogin
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}