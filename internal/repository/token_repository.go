package repository

import (
	"errors"
	"fmt"
	"gazes-auth/internal/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Sign generates a signed JWT token for the given user.
// It takes a user model as input and returns the signed token as a string.
// If an error occurs during the signing process, it is also returned.
func SignJWT(user model.User) (string, error) {
	claims := model.TokenPayload{
		ID:    user.ID,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			Subject:   "user",
			Issuer:    "admin",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	fmt.Println(os.Getenv("JWT_SECRET"))
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func VerifyJWT(tokenString string) (*model.TokenPayload, error) {
	claims := &model.TokenPayload{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
