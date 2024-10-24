package models

// URL represents a shortened URL entry
type URL struct {
    ID     uint   `gorm:"primaryKey"`
    Handle string `gorm:"uniqueIndex"`
    URL    string
}
