package handlers

import (
    "net/http"
    "go-shortlinks/models"
    "go-shortlinks/database"
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
    http.Redirect(w, r, url.URL, http.StatusFound)
}

