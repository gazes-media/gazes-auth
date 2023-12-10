package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}
	server.ListenAndServe()
}
