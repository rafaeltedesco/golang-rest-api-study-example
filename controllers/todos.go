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
