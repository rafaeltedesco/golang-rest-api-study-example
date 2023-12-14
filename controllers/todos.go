package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rafaeltedesco/rest-api/services"
)

type TodoDTO struct {
	Title       string `json:"title"`
	PlannedDate string `json:"plannedDate"`
}

var currDate = time.Now()

var todos = []services.Todo{
	{Id: 1, Title: "Create an ETL python project", PlannedDate: currDate.Add(2 * 24 * time.Hour)},
	{Id: 2, Title: "Create an API in Golang", PlannedDate: currDate.Add(5 * 24 * time.Hour)},
}

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

		newTodo := services.CreateTodo(&services.Todo{Title: Title, PlannedDate: parsedDate})
		jsonData, _ := json.Marshal(newTodo)
		w.Write(jsonData)
	}
}
