package main

import (
	"gazes-auth/database"
	"gazes-auth/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	database.Init()
}

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/login").HandlerFunc(routes.LoginHandler).Methods("POST")
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}

	server.ListenAndServe()
}
