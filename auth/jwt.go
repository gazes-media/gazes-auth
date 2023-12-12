package auth

import (
	"errors"
	"gazes-auth/database"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// UserClaim is the expected format of the JWT token, it contains the user ID and username
type UserClaim struct {
	jwt.RegisteredClaims
	Id       uint   `json:"id"`
	Email string `json:"email"`
}

// UserLogin is the expected format of the request body when logging in
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserRegister is the expected format of the request body when registering
type UserRegister struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SecretToken is the secret used to sign the JWT token, (it should be stored in an environment variable) #TODO: store in env var
var SecretToken = []byte("7c8mjYGG5H6VXyf6Zxqq6m69a2XNnPVC")

// Sign creates a JWT token from a user object
func Sign(user database.User) (string, error) {
	claims := UserClaim{
		Id:       user.ID,
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
