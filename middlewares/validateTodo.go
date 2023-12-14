package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	dtos "github.com/rafaeltedesco/rest-api/dtos/entities"
)

type ErrMessage struct {
	Message    string
	StatusCode int
}

func ValidateTodoMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todoDTO dtos.TodoInputDTO
		err := json.NewDecoder(r.Body).Decode(&todoDTO)

		if err != nil {
			fmt.Println("Invalid Body: ", err)
			errObj, _ := json.Marshal(ErrMessage{Message: "Invalid Body", StatusCode: http.StatusBadRequest})
			http.Error(w, string(errObj), http.StatusBadRequest)
			return
		}

		parsedDate, err := time.Parse("2006-01-02", todoDTO.PlannedDate)

		if err != nil {
			fmt.Println("Invalid date", err)
			errObj, _ := json.Marshal(ErrMessage{Message: "Error parsing date", StatusCode: http.StatusBadRequest})
			http.Error(w, string(errObj), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "parsedDate", parsedDate)
		ctx = context.WithValue(ctx, "Title", todoDTO.Title)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}
}
