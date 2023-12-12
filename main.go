package main

import (
	"gazes-auth/database"
	"gazes-auth/routes"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/jpfuentes2/go-env/autoload"
)

// init function is called before main function, so we can initialize our database connection here
func init() {
	database.Init()
	if os.Getenv("JWT_SECRET") == "" {
		panic("JWT_SECRET environment variable is not set")
	}

	if os.Getenv("PORT") == "" {
		panic("PORT environment variable is not set")
	}

	if os.Getenv("PASSWORD_SALT") == "" {
		panic("PASSWORD_SALT environment variable is not set")
	}
}

func main() {
	// Creating a new Mux Router
	r := mux.NewRouter()
	r.PathPrefix("/login").HandlerFunc(routes.LoginHandler).Methods("POST")
	r.PathPrefix("/register").HandlerFunc(routes.RegisterHandler).Methods("POST")
	r.PathPrefix("/user/@me").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		routes.GuardRoute(routes.UserMeHandler, w, r)
	}).Methods("GET")
	// Defining the Server
	server := &http.Server{
		Addr:    "0.0.0.0:" + os.Getenv("PORT"),
		Handler: r,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
