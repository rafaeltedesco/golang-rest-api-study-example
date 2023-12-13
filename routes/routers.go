package routes

import (
	"net/http"

	"github.com/rafaeltedesco/rest-api/controllers"
)

func InitHandlers() {
	http.HandleFunc("/", controllers.RedirectHome())
	http.HandleFunc("/healthy", controllers.ShowHealthy())
	http.HandleFunc("/todos", controllers.GetTodos())
}
