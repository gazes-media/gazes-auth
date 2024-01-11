package main

import (
	"gazes-auth/internal/api"
	"gazes-auth/internal/model"
	"gazes-auth/internal/utils"
	"log"
	"net/http"
	"os"
)

func init() {
	utils.ValidateEnvVars([]string{"JWT_SECRET", "PORT", "PASSWORD_SALT"})
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	router := api.NewRouter()

	log.Printf("Listening on port %s", os.Getenv("PORT"))

	if err := http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), router); err != nil {
		log.Fatalf(err.Error())
	}

	if err := utils.GetDB().AutoMigrate(&model.User{}); err != nil {
		log.Fatalf(err.Error())
	}
}
