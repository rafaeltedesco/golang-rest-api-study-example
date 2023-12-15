package mappers

import (
	"net/http"

	"github.com/rafaeltedesco/rest-api/dtos"
	"github.com/rafaeltedesco/rest-api/utils/customErrors"
)

func MapError(err error) dtos.ErrMessage {
	errorHttp := dtos.ErrMessage{Message: err.Error()}
	switch err {
	case customErrors.ErrNotFound:
		errorHttp.StatusCode = http.StatusNotFound
	case customErrors.ErrTaskIsDone:
		errorHttp.StatusCode = http.StatusBadRequest
	case customErrors.ErrCannotUndone:
		errorHttp.StatusCode = http.StatusBadRequest
	}
	return errorHttp
}
