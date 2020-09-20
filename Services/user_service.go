package service

import (
	"fmt"

	"github.com/ashishkhuraishy/bookstore_user_api/domain"
	"github.com/ashishkhuraishy/bookstore_user_api/utils/cryptoutils"
	"github.com/ashishkhuraishy/bookstore_user_api/utils/datetime"
	"github.com/ashishkhuraishy/bookstore_user_api/utils/errors"
)

// CreateUser inside service will handle all the business logic for creating a User
func CreateUser(user domain.User) (*domain.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = domain.StatusActive
	user.DateCreated = datetime.GetCurrentFormattedTime()
	user.Password = cryptoutils.HashPassword(user.Password)
	fmt.Println(user.Password)
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
	curentUser := &domain.User{ID: user.ID}
	if err := curentUser.Get(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName == "" {
			user.FirstName = curentUser.FirstName
		}
		if user.LastName == "" {
			user.LastName = curentUser.LastName
		}
		if user.Email == "" {
			user.Email = curentUser.Email
		}
	}
	user.DateCreated = curentUser.DateCreated
	user.DateUpdated = datetime.GetCurrentFormattedTime()

	if err := user.Update(isPartial); err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteUser : deletes a user with the given id / returns a restError
func DeleteUser(userID int64) *errors.RestError {
	user := &domain.User{ID: userID}

	if err := user.Delete(); err != nil {
		return err
	}

	return nil
}

// Search : returns all users with the status
func Search(status string) (domain.Users, *errors.RestError) {
	user := &domain.User{}

	return user.FindByStatus(status)
}
