package reqbind

import (
	"net/http"
)

func doValidate(target interface{}, r *http.Request) error {
	validatable, isValidatable := target.(Validatable)
	if !isValidatable {
		return nil
	}

	return validatable.Validate()
}

type Validatable interface {
	Validate() error
}
