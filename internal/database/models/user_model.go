package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`

	Auths []UserAuth
}

type UserAuth struct {
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Type     string `json:"type"` // can be google, password or discord

	UserID uint
}
