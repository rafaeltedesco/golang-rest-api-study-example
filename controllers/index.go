package controllers

import (
	"encoding/json"
	"net/http"
)

func RedirectHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/healthy", 301)
	}
}

func ShowHealthy() http.HandlerFunc {
	healthyMessage := make(map[string]bool)
	healthyMessage["healthy"] = true
	return func(w http.ResponseWriter, r *http.Request) {
		jsonData, _ := json.Marshal(healthyMessage)
		w.Write(jsonData)
	}
}
