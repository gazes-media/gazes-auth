package main

import (
	"gazes-auth/database"
	"gazes-auth/routes"
	"net/http"

	"github.com/gorilla/mux"
)

// init function is called before main function, so we can initialize our database connection here
func init() {
	database.Init()
}

func main() {
	// Creating a new Mux Router
	r := mux.NewRouter()
	// Handling Login Route
	r.PathPrefix("/login").HandlerFunc(routes.LoginHandler).Methods("POST")
	// Handling Register Route
	r.PathPrefix("/register").HandlerFunc(routes.RegisterHandler).Methods("POST")
	// Handling UserMe Route
	r.PathPrefix("/user/@me").HandlerFunc(routes.UserMeHandler).Methods("GET")
	// Defining the Server
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
