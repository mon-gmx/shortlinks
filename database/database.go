package database

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
)

var DB *gorm.DB // Global DB variable

// InitDB initializes the database connection
func InitDB(dsn string) error {
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    return err
}

