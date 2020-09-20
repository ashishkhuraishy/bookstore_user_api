package domain

import (
	"fmt"

	"github.com/ashishkhuraishy/bookstore_user_api/database/psql/userdb"
	"github.com/ashishkhuraishy/bookstore_user_api/utils/datetime"
	"github.com/ashishkhuraishy/bookstore_user_api/utils/errors"
	psqlerrors "github.com/ashishkhuraishy/bookstore_user_api/utils/psqlErrors"
)

const (
	queryInsertUser = `INSERT INTO users(first_name, last_name, email, password, status, date_created, date_updated) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id;`
	queryFindUser   = `SELECT first_name, last_name, email, status, date_created, date_updated FROM users WHERE id=$1`
	queryUpdateUser = `UPDATE users SET first_name=$1, last_name=$2, email=$3, date_updated=$4 WHERE id=$5`
)

var (
	users = make(map[int64]*User, 0)
)

// Save saves the user onto the database
func (u *User) Save() *errors.RestError {
	stmnt, err := userdb.Client.Prepare(queryInsertUser)
	defer stmnt.Close()
	if err != nil {
		return psqlerrors.ParseErrors(err)
	}

	u.DateCreated = datetime.GetCurrentFormattedTime()
	u.DateUpdated = u.DateCreated

	row := stmnt.QueryRow(
		u.FirstName,
		u.LastName,
		u.Email,
		u.Password,
		u.Status,
		u.DateCreated,
		u.DateUpdated,
	)

	if err := row.Err(); err != nil {
		return psqlerrors.ParseErrors(err)
	}

	if err := row.Scan(&u.ID); err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error while getting the id : %s", err.Error()))
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
		return errors.NewInternalServerError(err.Error())
	}

	row := stmnt.QueryRow(u.ID)

	if psqlErr := row.Err(); psqlErr != nil {
		return psqlerrors.ParseErrors(psqlErr)
	}

	if err = row.Scan(&u.FirstName, &u.LastName, &u.Email, &u.Status, &u.DateCreated, &u.DateUpdated); err != nil {
		return errors.NewNotFound("user not found")
	}

	return nil
}

// Update is used to update a  user in a database which returns a
// rest error
func (u *User) Update(isPartial bool) *errors.RestError {
	curentUser := &User{ID: u.ID}
	getError := curentUser.Get()
	if getError != nil {
		return getError
	}

	stmnt, err := userdb.Client.Prepare(queryUpdateUser)
	defer stmnt.Close()
	if err != nil {
		return errors.NewInternalServerError("Bad Query")
	}

	if isPartial {
		if u.FirstName == "" {
			u.FirstName = curentUser.FirstName
		}
		if u.LastName == "" {
			u.LastName = curentUser.LastName
		}
		if u.Email == "" {
			u.Email = curentUser.Email
		}
	}
	u.DateUpdated = datetime.GetCurrentFormattedTime()

	_, err = stmnt.Exec(u.FirstName, u.LastName, u.Email, u.DateUpdated)
	if err != nil {
		return psqlerrors.ParseErrors(err)
	}

	return nil
}
