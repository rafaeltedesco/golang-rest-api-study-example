package routes

import (
	"github.com/gorilla/mux"
	"github.com/rafaeltedesco/rest-api/controllers"
)

func HandleHealthy(r *mux.Router) {
	r.HandleFunc("/", controllers.RedirectHome()).Methods("GET")
	r.HandleFunc("/healthy", controllers.ShowHealthy()).Methods("GET")
}
