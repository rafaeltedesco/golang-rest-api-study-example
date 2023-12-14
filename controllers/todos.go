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
			return
		}
		err = services.DeleteTodo(id)
		if err != nil {
			httpError = dtos.ErrMessage{Message: err.Error(), StatusCode: http.StatusNotFound}
			jsonData, _ := json.Marshal(httpError)
			http.Error(w, string(jsonData), httpError.StatusCode)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func GetTodoById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			errorHttp := dtos.ErrMessage{Message: "Could not parse \"id\"", StatusCode: http.StatusBadRequest}
			http.Error(w, errorHttp.Message, errorHttp.StatusCode)
			return
		}
		todo, err := services.GetTodoById(id)

		if err != nil {
			errorHttp := dtos.ErrMessage{Message: err.Error(), StatusCode: http.StatusNotFound}
			http.Error(w, errorHttp.Message, errorHttp.StatusCode)
			return
		}

		jsonData, _ := json.Marshal(todo)

		w.Write(jsonData)
	}
}

func MarkTodoAsDone() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var err error
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			errorHttp := dtos.ErrMessage{Message: "Could not parse id", StatusCode: http.StatusBadRequest}
			w.WriteHeader(errorHttp.StatusCode)
			json.NewEncoder(w).Encode(errorHttp)
			return
		}
		err = services.MarkTodoAsDone(id)
		if err != nil {
			errorMessage := err.Error()
			errorHttp := dtos.ErrMessage{Message: err.Error()}
			switch errorMessage {
			case "Not found":
				errorHttp.StatusCode = http.StatusNotFound
			case "Already marked as done":
				errorHttp.StatusCode = http.StatusBadRequest
			case "Cannot undone a not finished task":
				errorHttp.StatusCode = http.StatusBadRequest
			}

			w.WriteHeader(errorHttp.StatusCode)
			json.NewEncoder(w).Encode(errorHttp)
			return
		}
		response := dtos.SuccessMessage{Message: "Task Marked as Done", StatusCode: http.StatusOK}
		json.NewEncoder(w).Encode(response)
	}
}

func UndoneTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var err error
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			errorHttp := dtos.ErrMessage{Message: "Could not parse id", StatusCode: http.StatusBadRequest}
			w.WriteHeader(errorHttp.StatusCode)
			json.NewEncoder(w).Encode(errorHttp)
			return
		}
		err = services.UndoneTodo(id)
		if err != nil {
			errorMessage := err.Error()
			errorHttp := dtos.ErrMessage{Message: err.Error()}
			switch errorMessage {
			case "Not found":
				errorHttp.StatusCode = http.StatusNotFound
			case "Already marked as done":
				errorHttp.StatusCode = http.StatusBadRequest
			case "Cannot undone a not finished task":
				errorHttp.StatusCode = http.StatusBadRequest
			}

			w.WriteHeader(errorHttp.StatusCode)
			json.NewEncoder(w).Encode(errorHttp)
			return
		}
		response := dtos.SuccessMessage{Message: "Undone Task successfully", StatusCode: http.StatusOK}
		json.NewEncoder(w).Encode(response)
	}
}
