package routes

import (
	"gazes-auth/database"
	"net/http"
)

type callback func(w http.ResponseWriter, r *http.Request, user *database.User)

func GuardRoute(cb callback, w http.ResponseWriter, r *http.Request) {
	user, ok := AuthGuard(w, r)
	if !ok {
		jsonErr := ErrorResponse{Error: "Error getting user"}
		jsonErr.Write(w, http.StatusInternalServerError)
		return
	}

	cb(w, r, user)
}
