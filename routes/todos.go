package routes

import (
	"github.com/gorilla/mux"
	"github.com/rafaeltedesco/rest-api/controllers"
	"github.com/rafaeltedesco/rest-api/middlewares"
)

func HandleTodos(r *mux.Router) {
	r.HandleFunc("/todos", controllers.GetTodos()).Methods("GET")
	r.HandleFunc("/todos", middlewares.ValidateTodoMiddleware(controllers.CreateTodo())).Methods("POST")
}