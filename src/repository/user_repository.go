package repository

import (
	"gazes-auth/src/model"
	"gazes-auth/src/utils"
)

// GetUserByID returns a user by their ID.
func GetUserByID(userID uint) (*model.User, error) {
	var user model.User
	if err := utils.GetDB().First(&user, userID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByEmail returns a user by their email address.
func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := utils.GetDB().First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateUser creates a new user in the database.
func CreateUser(user *model.User) error {
	if err := utils.GetDB().Create(user).Error; err != nil {
		return err
	}

	return nil
}
