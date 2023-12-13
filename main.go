package main

import (
	"fmt"
	"net/http"

	"github.com/rafaeltedesco/rest-api/routes"
)

func main() {
	address := ":8080"
	fmt.Printf("Server up and running on %s\n", address)
	routes.InitHandlers()
	http.ListenAndServe(address, nil)
}
