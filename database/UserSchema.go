package database

import (
	"gorm.io/gorm"
)

// UserSchema is the schema for the user table
type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Insert a new user into the database
func (u *User) Create() error {
	return DB.Create(u).Error
}

// Get a user by username
func (u *User) GetByUsername() error {
	return DB.Where("username = ?", u.Username).First(u).Error
}

// Get a user by email
func (u *User) GetByEmail() error {
	return DB.Where("email = ?", u.Email).First(u).Error
}

// Get a user by id
func (u *User) GetByID() error {
	return DB.Where("id = ?", u.ID).First(u).Error
}

// Update a user
func (u *User) Update() error {
	return DB.Save(u).Error
}

// Delete an user
func (u *User) Delete() error {
	return DB.Delete(u).Error
}
