package models

import (
	"time"
)

var currDate = time.Now()

type Todo struct {
	Id          int
	Title       string
	PlannedDate time.Time
	IsDone      bool
}

var todos = []Todo{
	{Id: 1, Title: "Create an ETL python project", PlannedDate: currDate.Add(2 * 24 * time.Hour)},
	{Id: 2, Title: "Create an API in Golang", PlannedDate: currDate.Add(5 * 24 * time.Hour)},
}

var nextId = 3

func GetTodos() []Todo {
	return todos
}

func CreateTodo(todo Todo) Todo {
	todo.Id = nextId
	todo.IsDone = false
	todos = append(todos, todo)
	nextId++
	return todo
}
