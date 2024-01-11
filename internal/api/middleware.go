package api

import (
	"context"
	"fmt"
	"gazes-auth/internal/repository"
	"gazes-auth/internal/utils"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			utils.RespondJSON(w, map[string]string{
				"error": "missing Authorization header",
			}, http.StatusBadRequest)
			return
		}

		if authHeader[:7] != "Bearer " {
			utils.RespondJSON(w, map[string]string{
				"error": "invalid Authorization header",
			}, http.StatusBadRequest)
			return
		}

		token := authHeader[7:]
		fmt.Println(token)

		payload, err := repository.VerifyJWT(token)

		if err != nil {
			utils.RespondJSON(w, map[string]string{
				"error":   "invalid token",
				"details": err.Error(),
			}, http.StatusBadRequest)
			return
		}

		user, err := repository.GetUserByID(payload.ID)

		if err != nil {
			utils.RespondJSON(w, map[string]string{
				"error": "user not found",
			}, http.StatusBadRequest)
			return
		}

		// Add the user to the context
		// This will be useful for the next handler
		// to retrieve the user from the context
		// and perform the necessary actions
		// (e.g. get the user's data
		ctx := context.WithValue(r.Context(), "user", user)

		// Call the next handler with the updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
