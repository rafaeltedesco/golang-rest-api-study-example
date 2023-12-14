package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rafaeltedesco/rest-api/dtos"
	"github.com/rafaeltedesco/rest-api/utils/validators"
)

func ValidateTodoMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todoDTO, err := validators.DecodeRequestBody(r)

		if err != nil {
			errObj, _ := json.Marshal(dtos.ErrMessage{Message: "Invalid Body", StatusCode: http.StatusBadRequest})
			http.Error(w, string(errObj), http.StatusBadRequest)
			return
		}

		parsedDate, err := validators.ParseDate(todoDTO.PlannedDate)

		if err != nil {
			fmt.Println("Invalid date", err)
			errObj, _ := json.Marshal(dtos.ErrMessage{Message: "Error parsing date", StatusCode: http.StatusBadRequest})
			http.Error(w, string(errObj), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "parsedDate", parsedDate)
		ctx = context.WithValue(ctx, "Title", todoDTO.Title)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}
}
