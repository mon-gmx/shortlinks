package handlers

import (
    "encoding/json"
    "net/http"
    "go-shortlinks/models"
    "go-shortlinks/database"
)

// ShortenURL handles the URL shortening
func ShortenURL(w http.ResponseWriter, r *http.Request) {
    var url models.URL
    if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    if err := database.DB.Create(&url).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusOK)
}
