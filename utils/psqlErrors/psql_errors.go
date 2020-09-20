package psqlerrors

import (
	"fmt"
	"strings"

	"github.com/ashishkhuraishy/bookstore_user_api/utils/errors"
	"github.com/lib/pq"
)

const (
	duplicateValue = "23505"
	notFound       = "no rows in result set"
)

// ParseErrors takes in an error obj and converts it into a
// psql error and return appropriate rest error
func ParseErrors(err error) *errors.RestError {
	psqlError, ok := err.(*pq.Error)
	if !ok {
		if strings.Contains(err.Error(), notFound) {
			return errors.NewNotFound("no record matching the given id")
		}
		return errors.NewInternalServerError(fmt.Sprintf("database error -> %s", err.Error()))
	}
	switch psqlError.Code {
	case duplicateValue:
		return errors.NewBadRequest("data already in use")
	default:
		return errors.NewInternalServerError(fmt.Sprintf("Code : %s Message : %s", psqlError.Code, psqlError.Message))
	}
}
