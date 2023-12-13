package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rafaeltedesco/rest-api/controllers"
)

func InitHandlers() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.RedirectHome()).Methods("GET")
	r.HandleFunc("/healthy", controllers.ShowHealthy()).Methods("GET")
	r.HandleFunc("/todos", controllers.GetTodos()).Methods("GET")
	r.HandleFunc("/todos", controllers.CreateTodo()).Methods("POST")
	http.Handle("/", r)
}
