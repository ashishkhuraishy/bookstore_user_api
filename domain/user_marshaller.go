package domain

import (
	"encoding/json"

	"github.com/ashishkhuraishy/bookstore_user_api/utils/errors"
)

// PublicUser is the interface returned to public req
type PublicUser struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// PrivateUser is the interface returned to private req
type PrivateUser struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
}

// Marshaller is used to marshell the user before sending to
// as a response
func (u *User) Marshaller(isPublic bool) interface{} {
	userJSON, err := json.Marshal(u)
	if err != nil {
		return errors.NewInternalServerError("could not marshall the data")
	}

	if isPublic {
		var user PublicUser
		json.Unmarshal(userJSON, &user)
		return user
	}

	var user PrivateUser
	json.Unmarshal(userJSON, &user)
	return user

}

// Marshaller is used to marshell the user before sending to
// as a response
func (users Users) Marshaller(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for i, user := range users {
		result[i] = user.Marshaller(isPublic)
	}

	return result
}
