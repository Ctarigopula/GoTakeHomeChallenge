package coordinators

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"take-home-challenge/models"
)

type UsersCoordinator interface {
	Create(user *models.User) error
	Read(userID int) (*models.User, error)
}

type usersCoordinator struct {
	db *gorm.DB
}

func NewUsersCoordinator(db *gorm.DB) UsersCoordinator {
	return usersCoordinator{
		db: db,
	}
}

func (u usersCoordinator) Create(user *models.User) error {
	if user.ClientID == 0 {
		err := errors.Wrap(&validationError{message: "ClientID is invalid."}, "Failed to create user")
		return err
	}

	if user.FirstName == "" {
		err := errors.Wrap(&validationError{message: "FirstName is invalid."}, "Failed to create user")
		return err
	}
	if user.LastName == "" {
		err := errors.Wrap(&validationError{message: "LastName is invalid."}, "Failed to create user")
		return err
	}

	if err := u.db.Table("client_users").Create(user).Error; err != nil {
		return errors.Wrap(err, "failed to create new client user in coordinator")
	}

	fmt.Println("New user!")

	return nil
}

func (u usersCoordinator) Read(userID int) (*models.User, error) {
	query := u.db.Table("client_users").Where("id = ?", userID)
	user := &models.User{
		ID: userID,
	}
	if err := query.First(user).Error; err != nil {
		return nil, errors.Wrap(err, "failed to read user in coordinator")
	}
	fmt.Println("Found the user!")
	return user, nil
}
