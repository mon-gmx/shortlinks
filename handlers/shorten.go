package handlers

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"go-shortlinks/database"
	"go-shortlinks/models"
	"net/http"

	"gorm.io/gorm"
)

// ShortenURL handles the URL shortening
func ShortenURL(w http.ResponseWriter, r *http.Request) {
	action := "insert"
	if r.Method == http.MethodPut {
		action = "update"
	} else if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var url models.URL
	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Encode the URL to base64 to avoid storing it in plain text
	url.URL = base64.StdEncoding.EncodeToString([]byte(url.URL))

	var existingURL models.URL
	if action == "update" {
		// Check if the entry exists by handle
		if err := database.DB.Where("handle = ?", url.Handle).First(&existingURL).Error; err == nil {
			// If the entry exists, update it
			if err := database.DB.Model(&existingURL).Updates(url).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			// If the entry is not found and we want to insert a new one
			if err := database.DB.Create(&url).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// If action is insert (POST)
		if err := database.DB.Create(&url).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "shortlink was stored successfully")
}
