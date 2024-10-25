package handlers

import (
    "encoding/json"
    "net/http"
    "go-shortlinks/models"
    "go-shortlinks/database"
)

func GetAllURLs(w http.ResponseWriter, r *http.Request) {
    var urls []models.URL

    if result := database.DB.Find(&urls); result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(urls)
}
