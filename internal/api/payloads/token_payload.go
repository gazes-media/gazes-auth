package payloads

import "github.com/golang-jwt/jwt/v5"

type TokenPayload struct {
	jwt.RegisteredClaims
	ID       uint   `json:"id"`
	Username string `json:"username"`
}
