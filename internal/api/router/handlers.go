package router

import (
	"encoding/json"
	"gazes-auth/internal/api/payloads"
	"gazes-auth/internal/database/models"
	"gazes-auth/internal/database/repositories"
	"gazes-auth/pkg/utils"
	"net/http"

	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

// PostLogin handles the HTTP POST request for user login.
// It decodes the JSON request body into a User struct, retrieves the user from the database based on the email,
// hashes the password, and checks if the email and password match.
// If successful, it generates a JWT token and responds with the token.
// If any error occurs during the process, it responds with an error message.
func PostLogin(w http.ResponseWriter, r *http.Request) {

	var userLogin payloads.UserLogin
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

	user, err := repositories.GetUserByEmail(userLogin.Email)
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

	token, err := utils.SignJWT(*user)

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

	// décoder la payload
	var registerPayload payloads.UserRegister
	if err := json.NewDecoder(r.Body).Decode(&registerPayload); err != nil {
		utils.RespondJSON(w, map[string]string{
			"error": "invalid request body",
		}, http.StatusBadRequest)

		return
	}

	// valider les inputs
	if err := validator.Validate(registerPayload); err != nil {
		json.Unmarshal([]byte(err.Error()), &err)
		utils.RespondJSON(w, map[string]interface{}{"error": err}, http.StatusBadRequest)

		return
	}

	// check if the user already have an account with this email

	var user *models.User

	utils.GetDB().Preload("Auths", func(db *gorm.DB) *gorm.DB {
		return db.Where("type = ?", "password").Where("password = ?", utils.HashPassword(registerPayload.Password))
	}).First(user)

	if user != nil {
		utils.RespondJSON(w, map[string]string{
			"error": "user with this email already exists",
		}, http.StatusBadRequest)
	}
}

// GetMe handles the HTTP GET request for retrieving the user's data.
// It retrieves the user from the context and responds with the user's data.
func GetMe(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*models.User)
	utils.RespondJSON(w, user, http.StatusOK)
}
