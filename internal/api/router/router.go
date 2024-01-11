package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/login", PostLogin).Methods("POST")
	router.HandleFunc("/register", PostRegister).Methods("POST")
	router.Handle("/@me", AuthMiddleware(http.HandlerFunc(GetMe))).Methods("GET")

	return router
}
