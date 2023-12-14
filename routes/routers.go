package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitHandlers() {
	r := mux.NewRouter()
	HandleHealthy(r)
	HandleTodos(r)
	http.Handle("/", r)
}
