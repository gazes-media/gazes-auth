package model

import "github.com/golang-jwt/jwt/v5"

type TokenPayload struct {
	jwt.RegisteredClaims
	ID    uint   `json:"id"`    // ID is the user's ID.
	Email string `json:"email"` // Email is the user's email address.
}
