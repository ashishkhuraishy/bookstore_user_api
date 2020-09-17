package domain

import (
	"strings"

	"github.com/ashishkhuraishy/bookstore_user_api/utils/errors"
)

// User domain
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	DateTime  string `json:"datetime"`
}

// Validate : Used to validate a given user
func (u *User) Validate() *errors.RestError {
	if u.ID < 1 {
		return errors.NewBadRequest("invalid user id")
	}
	u.Email = strings.TrimSpace(u.Email)
	if u.Email == "" {
		return errors.NewBadRequest("invalid email")
	}
	return nil
}
