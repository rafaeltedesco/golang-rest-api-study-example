package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/rafaeltedesco/rest-api/dtos"
	"github.com/rafaeltedesco/rest-api/models"
	"github.com/rafaeltedesco/rest-api/services"
)

func GetTodos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonData, _ := json.Marshal(services.GetTodos())
		w.Write(jsonData)
	}
}

func CreateTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parsedDate := r.Context().Value("parsedDate").(time.Time)
		Title := r.Context().Value("Title").(string)

		newTodo := services.CreateTodo(models.Todo{Title: Title, PlannedDate: parsedDate})
		jsonData, _ := json.Marshal(newTodo)
		w.Write(jsonData)
	}
}

func DeleteTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var err error
		var httpError dtos.ErrMessage
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			httpError = dtos.ErrMessage{Message: "Could not parse id!", StatusCode: http.StatusBadRequest}
			http.Error(
				w,
				httpError.Message,
				httpError.StatusCode,
			)
		}
		err = services.DeleteTodo(id)
		if err != nil {
			httpError = dtos.ErrMessage{Message: err.Error(), StatusCode: http.StatusNotFound}
			jsonData, _ := json.Marshal(httpError)
			http.Error(w, string(jsonData), httpError.StatusCode)
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
