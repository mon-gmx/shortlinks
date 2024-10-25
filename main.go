package main

import (
    "log"
    "fmt"
    "net/http"
    "go-shortlinks/config"
    "go-shortlinks/database"
    "go-shortlinks/models"
    "go-shortlinks/handlers"
)

func main() {

    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("failed to load config: %v", err)
    }

    // Build DSN for PostgreSQL (prod)
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
        cfg.Database.Prod.Host,
        cfg.Database.Prod.User,
        cfg.Database.Prod.Password,
        cfg.Database.Prod.DBName,
        cfg.Database.Prod.Port,
        cfg.Database.Prod.SSLMode,
        cfg.Database.Prod.TimeZone,
    )
    
    if err := database.InitDB(dsn); err != nil {
        log.Fatalf("failed to connect to PostgreSQL: %v", err)
    }

    // Migrate the schema
    database.DB.AutoMigrate(&models.URL{})

    // Set up routes
    http.HandleFunc("/shorten", handlers.ShortenURL)
    http.HandleFunc("/", handlers.RedirectURL)
    http.HandleFunc("/urls", handlers.GetAllURLs)

    log.Println("Server is running on :8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}

