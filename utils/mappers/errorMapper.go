package mappers

import (
	"net/http"

	"github.com/rafaeltedesco/rest-api/dtos"
)

func MapError(err error) dtos.ErrMessage {
	errorHttp := dtos.ErrMessage{Message: err.Error()}
	switch err.Error() {
	case "Not found":
		errorHttp.StatusCode = http.StatusNotFound
	case "Already marked as done":
		errorHttp.StatusCode = http.StatusBadRequest
	case "Cannot undone a not finished task":
		errorHttp.StatusCode = http.StatusBadRequest
	}
	return errorHttp
}
