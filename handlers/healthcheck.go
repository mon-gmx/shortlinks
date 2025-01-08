package handlers

import (
	"log"
	"net/http"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {

	if !IsValidMethod(w, r, http.MethodGet) {
		return
	}
	var urls []models.URL
	w.Header().Set("Content-Type", "application/json")
}
