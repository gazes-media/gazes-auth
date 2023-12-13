package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`            // Username is the user's username.
	Email    string `gorm:"unique" json:"email"` // Email is the user's email address.
	Password string `json:"password"`            // Password is the user's password.
}

type UserLogin struct {
	Email    string `json:"email" validate:"regexp=^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$"`
	Password string `json:"password" validate:"min=8"`
}

type UserRegister struct {
	Email    string `json:"email" validate:"regexp=^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$"`
	Password string `json:"password" validate:"min=8"`
	Username string `json:"username" validate:"min=3,max=40,regexp=^[a-zA-Z0-9_]+$"`
}
