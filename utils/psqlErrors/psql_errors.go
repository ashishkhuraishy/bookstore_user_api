package psqlerrors

import (
	"fmt"

	"github.com/ashishkhuraishy/bookstore_user_api/utils/errors"
	"github.com/lib/pq"
)

const (
	duplicateValue = "23505"
)

// ParseErrors takes in an error obj and converts it into a
// psql error and return appropriate rest error
func ParseErrors(err error) *errors.RestError {
	psqlError := err.(*pq.Error)
	switch psqlError.Code {
	case duplicateValue:
		return errors.NewBadRequest("data already exists")
	default:
		return errors.NewInternalServerError(fmt.Sprintf("Code : %s Message : %s", psqlError.Code, psqlError.Message))
	}
}
