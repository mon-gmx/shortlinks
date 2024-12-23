package handlers

import (
	"encoding/base64"
	"go-shortlinks/database"
	"go-shortlinks/models"
	"net/http"
)

// RedirectURL handles URL redirection based on the handle
func RedirectURL(w http.ResponseWriter, r *http.Request) {
	if !IsValidMethod(w, r, http.MethodGet) {
		return
	}
	handle := r.URL.Path[1:] // Remove leading "/"

	var url models.URL
	if err := database.DB.Where("handle = ?", handle).First(&url).Error; err != nil {
		http.NotFound(w, r)
		return
	}
	decodedURL, err := base64.StdEncoding.DecodeString(url.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, string(decodedURL), http.StatusFound)
}
