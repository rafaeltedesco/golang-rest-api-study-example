package controllers

import (
	"encoding/json"
	"fmt"
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
		var todoDTO TodoDTO
		err := json.NewDecoder(r.Body).Decode(&todoDTO)

		if err != nil {
			fmt.Println("Invalid Body: ", err)
			http.Error(w, "Invalid Body", http.StatusBadRequest)
		}

		fmt.Println(todoDTO)

		parsedData, err := time.Parse("2006-01-02", todoDTO.PlannedDate)

		if err != nil {
			fmt.Println("Invalid date", err)
			http.Error(w, "Error parsing date", http.StatusBadRequest)
		}

		newTodo := Todo{Id: len(todos) + 1, Title: todoDTO.Title, PlannedDate: parsedData}

		todos = append(todos, newTodo)
		jsonData, _ := json.Marshal(newTodo)
		w.Write(jsonData)
	}
}
