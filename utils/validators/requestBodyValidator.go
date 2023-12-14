package validators

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rafaeltedesco/rest-api/dtos"
)

func DecodeRequestBody(r *http.Request) (dtos.TodoInputDTO, error) {
	var todoDTO dtos.TodoInputDTO
	err := json.NewDecoder(r.Body).Decode(&todoDTO)
	return todoDTO, err
}

func ParseDate(plannedDate string) (time.Time, error) {
	parsedDate, err := time.Parse("2006-01-02", plannedDate)
	return parsedDate, err
}
