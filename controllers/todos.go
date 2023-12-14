package controllers

import (
	"encoding/json"
	"net/http"
	"time"

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
