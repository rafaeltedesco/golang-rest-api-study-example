package controllers

import (
	"encoding/json"
	"net/http"
	"time"
)

type Todo struct {
	Id          int
	Title       string
	PlannedDate time.Time
	IsDone      bool
}

type TodoDTO struct {
	Title       string `json:"title"`
	PlannedDate string `json:"plannedDate"`
}

var currDate = time.Now()

var todos = []Todo{
	{Id: 1, Title: "Create an ETL python project", PlannedDate: currDate.Add(2 * 24 * time.Hour)},
	{Id: 2, Title: "Create an API in Golang", PlannedDate: currDate.Add(5 * 24 * time.Hour)},
}

func GetTodos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonData, _ := json.Marshal(todos)
		w.Write(jsonData)
	}
}

func CreateTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parsedDate := r.Context().Value("parsedDate").(time.Time)
		Title := r.Context().Value("Title").(string)

		newTodo := Todo{Id: len(todos) + 1, Title: Title, PlannedDate: parsedDate}

		todos = append(todos, newTodo)
		jsonData, _ := json.Marshal(newTodo)
		w.Write(jsonData)
	}
}
