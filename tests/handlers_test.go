package handlers_test

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "go-shortlinks/handlers"
    "go-shortlinks/models"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "go-shortlinks/database"
    "go-shortlinks/config"
)

func setupTestDB() (*gorm.DB, error) {
    cfg, err := config.LoadConfig()
    if err != nil {
        return nil, err
    }

    db, err := gorm.Open(sqlite.Open(cfg.Database.Test.DSN), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    db.AutoMigrate(&models.URL{})
    database.DB = db // Set the global DB variable for testing

    return db, nil
}

func TestShortenURL(t *testing.T) {
    _, err := setupTestDB()
    if err != nil {
        t.Fatal(err)
    }

    body := `{"handle":"testhandle", "url":"http://example.com"}`
    req, err := http.NewRequest("POST", "/shorten", strings.NewReader(body))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json") // Set the content type to JSON

    rr := httptest.NewRecorder()
    handlers.ShortenURL(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }
}

func TestRedirectURL(t *testing.T) {
    db, err := setupTestDB()
    if err != nil {
        t.Fatal(err)
    }

    // Create a URL entry
    db.Create(&models.URL{Handle: "testhandle", URL: "http://example.com"})

    req, err := http.NewRequest("GET", "/testhandle", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handlers.RedirectURL(rr, req)

    if status := rr.Code; status != http.StatusFound {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusFound)
    }
}

