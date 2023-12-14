package routes

import (
	"github.com/gorilla/mux"
	"github.com/rafaeltedesco/rest-api/controllers"
	"github.com/rafaeltedesco/rest-api/middlewares"
)

func HandleTodos(r *mux.Router) {
	todosRouter := r.PathPrefix("/todos").Subrouter()
	todosRouter.HandleFunc("/", controllers.GetTodos()).Methods("GET")
	todosRouter.HandleFunc("/{id}", controllers.GetTodoById()).Methods("GET")
	todosRouter.HandleFunc("/", middlewares.ValidateTodoMiddleware(controllers.CreateTodo())).Methods("POST")
	todosRouter.HandleFunc("/{id}/finish", controllers.MarkTodoAsDone()).Methods("PATCH")
	todosRouter.HandleFunc("/{id}", controllers.DeleteTodo()).Methods("DELETE")
}
