package domain

import (
	"fmt"

	"github.com/ashishkhuraishy/bookstore_user_api/database/psql/userdb"
	"github.com/ashishkhuraishy/bookstore_user_api/utils/errors"
	psqlerrors "github.com/ashishkhuraishy/bookstore_user_api/utils/psqlErrors"
)

const (
	// StatusActive global status active key
	StatusActive = "active"

	queryInsertUser   = `INSERT INTO users(first_name, last_name, email, password, status, date_created, date_updated) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id;`
	queryFindUser     = `SELECT first_name, last_name, email, status, date_created, date_updated FROM users WHERE id=$1`
	queryUpdateUser   = `UPDATE users SET first_name=$1, last_name=$2, email=$3, date_updated=$4 WHERE id=$5`
	queryDeleteUser   = `DELETE FROM users WHERE id=$1`
	queryGetAllActive = `SELECT id, first_name, last_name, email, status, date_created, date_updated FROM users WHERE status=$1`
)

// Save saves the user onto the database
func (u *User) Save() *errors.RestError {
	stmnt, err := userdb.Client.Prepare(queryInsertUser)
	defer stmnt.Close()
	if err != nil {
		return psqlerrors.ParseErrors(err)
	}

	err = stmnt.QueryRow(
		u.FirstName,
		u.LastName,
		u.Email,
		u.Password,
		u.Status,
		u.DateCreated,
		u.DateUpdated,
	).Scan(&u.ID)

	if err != nil {
		return psqlerrors.ParseErrors(err)
	}

	return nil
}

// Get returns a user not found error or gets the user and
// save it to the obj
func (u *User) Get() *errors.RestError {
	if u.ID < 1 {
		return errors.NewBadRequest("invalid user id")
	}

	stmnt, err := userdb.Client.Prepare(queryFindUser)
	defer stmnt.Close()
	if err != nil {
		return psqlerrors.ParseErrors(err)
	}

	err = stmnt.QueryRow(u.ID).Scan(
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Status,
		&u.DateCreated,
		&u.DateUpdated,
	)

	if err != nil {
		return psqlerrors.ParseErrors(err)
	}

	return nil
}

// Update is used to update a  user in a database which returns a
// rest error
func (u *User) Update(isPartial bool) *errors.RestError {
	stmnt, err := userdb.Client.Prepare(queryUpdateUser)
	defer stmnt.Close()
	if err != nil {
		return errors.NewInternalServerError("Bad Query")
	}

	_, err = stmnt.Exec(u.FirstName, u.LastName, u.Email, u.DateUpdated, u.ID)
	if err != nil {
		return psqlerrors.ParseErrors(err)
	}

	return nil
}

// Delete : To remove a user use this function which returns
// a RestError
func (u *User) Delete() *errors.RestError {
	stmnt, err := userdb.Client.Prepare(queryDeleteUser)
	defer stmnt.Close()
	if err != nil {
		return errors.NewInternalServerError("invalid query")
	}

	_, err = stmnt.Exec(u.ID)
	if err != nil {
		return psqlerrors.ParseErrors(err)
	}

	return nil
}

// FindByStatus : To return all active users use this function which returns
// a RestError and list of users
func (u *User) FindByStatus(status string) ([]User, *errors.RestError) {
	stmnt, err := userdb.Client.Prepare(queryGetAllActive)
	defer stmnt.Close()
	if err != nil {
		return nil, errors.NewInternalServerError("invalid query")
	}

	rows, err := stmnt.Query(status)
	defer rows.Close()
	if err != nil {
		return nil, psqlerrors.ParseErrors(err)
	}

	users := make([]User, 0)
	for rows.Next() {
		var user User
		err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Status,
			&user.DateCreated,
			&user.DateUpdated,
		)

		if err != nil {
			return nil, psqlerrors.ParseErrors(err)
		}

		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, errors.NewNotFound(fmt.Sprintf("no users having status %s", status))
	}

	return users, nil
}
