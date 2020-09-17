package domain

import (
	"fmt"

	"github.com/ashishkhuraishy/bookstore_user_api/utils/errors"
)

var (
	users = make(map[int64]*User, 0)
)

// Save saves the user onto the database
func (u *User) Save() *errors.RestError {
	if users[u.ID] != nil {
		return errors.NewBadRequest(fmt.Sprintf("User %d already exists", u.ID))
	}

	users[u.ID] = u
	return nil
}

// Get returns a user not found error or gets the user and
// save it to the obj
func (u *User) Get() *errors.RestError {
	if u.ID < 1 {
		return errors.NewBadRequest("invalid user id")
	}
	user := users[u.ID]
	if user == nil {
		return errors.NewNotFound(fmt.Sprintf("could not find user %d", u.ID))
	}

	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Email = user.Email
	u.DateTime = user.DateTime
	return nil
}
