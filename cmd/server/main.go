package main

import (
	"fmt"
	"gazes-auth/internal/api/router"
	"gazes-auth/internal/database/models"
	"gazes-auth/pkg/utils"
	"log"
	"net/http"
)

func main() {
	router := router.NewRouter()

	vars, err := utils.Getenvars([]string{"PORT"})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Listening on port %s...\n", vars["PORT"])
	http.ListenAndServe("0.0.0.0:"+vars["PORT"], router)

	utils.GetDB().AutoMigrate(&models.User{})
}
