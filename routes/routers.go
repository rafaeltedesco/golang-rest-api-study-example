package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rafaeltedesco/rest-api/controllers"
	"github.com/rafaeltedesco/rest-api/middlewares"
)

func InitHandlers() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.RedirectHome()).Methods("GET")
	r.HandleFunc("/healthy", controllers.ShowHealthy()).Methods("GET")
	r.HandleFunc("/todos", controllers.GetTodos()).Methods("GET")
	r.HandleFunc("/todos", middlewares.ValidateTodoMiddleware(controllers.CreateTodo())).Methods("POST")
	http.Handle("/", r)
}
