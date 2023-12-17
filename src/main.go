package main

import (
	"fmt"
	"gazes-auth/src/api"
	"gazes-auth/src/model"
	"gazes-auth/src/utils"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %s", err)
	}

	utils.ValidateEnvVars([]string{"JWT_SECRET", "PORT", "PASSWORD_SALT"})

}

func main() {
	router := api.NewRouter()

	fmt.Printf("Listening on port %s...\n", os.Getenv("PORT"))
	http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), router)

	utils.GetDB().AutoMigrate(&model.User{})
}
