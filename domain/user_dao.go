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
