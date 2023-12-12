package api

import (
	"encoding/json"
	"fmt"
	"gazes-auth/src/model"
	"gazes-auth/src/repository"
	"gazes-auth/src/utils"
	"net/http"

	"gopkg.in/validator.v2"
)

// PostLogin handles the HTTP POST request for user login.
// It decodes the JSON request body into a User struct, retrieves the user from the database based on the email,
// hashes the password, and checks if the email and password match.
// If successful, it generates a JWT token and responds with the token.
// If any error occurs during the process, it responds with an error message.
func PostLogin(w http.ResponseWriter, r *http.Request) {

	var userLogin model.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {
		utils.RespondJSON(w, map[string]string{
			"error": "invalid request body",
		}, http.StatusBadRequest)
		return
	}

	if err := validator.Validate(userLogin); err != nil {
		json.Unmarshal([]byte(err.Error()), &err)
		utils.RespondJSON(w, map[string]interface{}{"error": err}, http.StatusBadRequest)

		return
	}

	user, err := repository.GetUserByEmail(userLogin.Email)
	if err != nil {
		utils.RespondJSON(w, map[string]string{
			"error": "invalid credentials",
		}, http.StatusBadRequest)
		return
	}

	hashedPassword := utils.HashPassword(userLogin.Password)

	if user.Password != hashedPassword {
		utils.RespondJSON(w, map[string]string{
			"error": "invalid credentials",
		}, http.StatusBadRequest)
		return
	}

	token, err := repository.SignJWT(*user)

	if err != nil {
		utils.RespondJSON(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: token,
	})

}

// PostRegister handles the HTTP POST request for user registration.
// It decodes the JSON request body into a User struct, checks if the email already exists in the database,
// hashes the password, and creates the user in the database.
// If successful, it generates a JWT token and responds with the token.
// If any error occurs during the process, it responds with an error message.
func PostRegister(w http.ResponseWriter, r *http.Request) {

	var userRegister model.UserRegister
	if err := json.NewDecoder(r.Body).Decode(&userRegister); err != nil {
		utils.RespondJSON(w, map[string]string{
			"error": "invalid request body",
		}, http.StatusBadRequest)
		return
	}

	fmt.Println(userRegister)

	if err := validator.Validate(userRegister); err != nil {
		json.Unmarshal([]byte(err.Error()), &err)
		utils.RespondJSON(w, map[string]interface{}{"error": err}, http.StatusBadRequest)

		return
	}

	_, err := repository.GetUserByEmail(userRegister.Email)
	if err == nil {
		utils.RespondJSON(w, map[string]string{
			"error": "email already exists",
		}, http.StatusBadRequest)
		return
	}

	user := model.User{
		Email:    userRegister.Email,
		Password: utils.HashPassword(userRegister.Password),
		Username: userRegister.Username,
	}

	if err := repository.CreateUser(&user); err != nil {
		utils.RespondJSON(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	token, err := repository.SignJWT(user)

	if err != nil {
		utils.RespondJSON(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: token,
	})
}

// GetMe handles the HTTP GET request for retrieving the user's data.
// It retrieves the user from the context and responds with the user's data.
func GetMe(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*model.User)
	utils.RespondJSON(w, user, http.StatusOK)
}
