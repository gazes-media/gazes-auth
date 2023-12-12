package auth

import (
	"errors"
	"gazes-auth/database"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// UserClaim is the expected format of the JWT token, it contains the user ID and email
type UserClaim struct {
	jwt.RegisteredClaims
	Id    uint   `json:"id"`
	Email string `json:"email"`
}

// UserLogin is the expected format of the request body when logging in
type UserLogin struct {
	Email    string `json:"email" validate:"regexp=^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$"`
	Password string `json:"password" validate:"min=8"`
}

// UserRegister is the expected format of the request body when registering
type UserRegister struct {
	Username string `json:"username" validate:"min=3,max=40,regexp=^[a-zA-Z0-9_]+$"`
	Email    string `json:"email" validate:"regexp=^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$"`
	Password string `json:"password" validate:"min=8"`
}

var SecretToken = []byte(os.Getenv("JWT_SECRET"))

// Sign creates a JWT token from a user object
func Sign(user database.User) (string, error) {
	claims := UserClaim{
		Id:    user.ID,
		Email: user.Email,
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

// Verify checks if a JWT token is valid and returns the user Claim Object
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
