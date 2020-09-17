package domain

import (
	"strings"

	"github.com/ashishkhuraishy/bookstore_user_api/utils/errors"
)

// User domain
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated,omitempty"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

// Validate : Used to validate a given user
func (u *User) Validate() *errors.RestError {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.NewBadRequest("invalid email")
	}
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	return nil
}
