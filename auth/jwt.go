package auth

import (
	"errors"
	"gazes-auth/database"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWT is a wrapper around the jwt.Token type

type UserClaim struct {
	jwt.RegisteredClaims
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegister struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var SecretToken = []byte("7c8mjYGG5H6VXyf6Zxqq6m69a2XNnPVC")

func Sign(user database.User) (string, error) {
	claims := UserClaim{
		Id:       int(user.ID),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			Subject:   "user",
			Issuer:    "admin",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretToken)
}

func Verify(tokenString string) (*UserClaim, error) {
	claims := &UserClaim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return SecretToken, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
