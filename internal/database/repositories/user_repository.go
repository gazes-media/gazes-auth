package repositories

import (
	"gazes-auth/internal/database/models"
	"gazes-auth/pkg/utils"
)

// GetUserByID returns a user by their ID.
func GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	if err := utils.GetDB().First(&user, userID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByEmail returns a user by their email address.
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := utils.GetDB().First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateUser creates a new user in the database.
func CreateUser(user *models.User) error {
	if err := utils.GetDB().Create(user).Error; err != nil {
		return err
	}

	return nil
}
