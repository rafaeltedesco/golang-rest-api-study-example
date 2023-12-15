package models

import (
	"time"

	"github.com/rafaeltedesco/rest-api/utils/customErrors"
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

func DeleteTodo(id int) error {
	for idx, todo := range todos {
		if todo.Id == id {
			todos = append(todos[:idx], todos[idx+1:]...)
			return nil
		}
	}
	return customErrors.ErrNotFound
}

func GetTodoById(id int) (*Todo, error) {
	for idx, todo := range todos {
		if todo.Id == id {
			return &todos[idx], nil
		}
	}
	return &Todo{}, customErrors.ErrNotFound
}

func MarkTodoAsDone(id int) error {
	todo, err := GetTodoById(id)
	if err != nil {
		return err
	}
	if todo.IsDone {
		return customErrors.ErrTaskIsDone
	}
	todo.IsDone = true
	return nil
}

func UndoneTodo(id int) error {
	todo, err := GetTodoById(id)
	if err != nil {
		return err
	}
	if !todo.IsDone {
		return customErrors.ErrCannotUndone
	}
	todo.IsDone = false
	return nil
}
