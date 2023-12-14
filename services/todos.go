package services

import "github.com/rafaeltedesco/rest-api/models"

func GetTodos() []models.Todo {
	return models.GetTodos()
}

func CreateTodo(todo models.Todo) models.Todo {
	newTodo := models.CreateTodo(todo)
	return newTodo
}
