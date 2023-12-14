package services

import (
	"github.com/rafaeltedesco/rest-api/dtos"
	"github.com/rafaeltedesco/rest-api/models"
)

func GetTodos() []dtos.TodoOutputDTO {
	todos := models.GetTodos()
	var output []dtos.TodoOutputDTO
	for _, todo := range todos {
		output = append(output, dtos.TodoOutputDTO{
			Id:          todo.Id,
			Title:       todo.Title,
			PlannedDate: todo.PlannedDate,
			IsDone:      todo.IsDone,
		})
	}
	return output
}

func CreateTodo(todo models.Todo) dtos.TodoOutputDTO {
	newTodo := models.CreateTodo(todo)
	output := dtos.TodoOutputDTO{Id: newTodo.Id, Title: newTodo.Title, PlannedDate: newTodo.PlannedDate, IsDone: newTodo.IsDone}
	return output
}

func DeleteTodo(id int) error {
	return models.DeleteTodo(id)
}

func GetTodoById(id int) (models.Todo, error) {
	return models.GetTodoById(id)
}
