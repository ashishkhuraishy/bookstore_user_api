package service

import (
	"github.com/ashishkhuraishy/bookstore_user_api/domain"
	"github.com/ashishkhuraishy/bookstore_user_api/utils/errors"
)

// CreateUser inside service will handle all the business logic for creating a User
func CreateUser(user domain.User) (*domain.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUser : gets a user with the given id / returns a resterror
func GetUser(userID int64) (*domain.User, *errors.RestError) {
	user := domain.User{ID: userID}
	err := user.Get()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser : updates a user with the given id / returns a restError
func UpdateUser(user domain.User, isPartial bool) (*domain.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Update(isPartial); err != nil {
		return nil, err
	}

	return &user, nil
}
