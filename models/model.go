package models

import "time"

type User struct {
	UserID    uint      `gorm:"primaryKey"`
	Username  string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(255);not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
}

type Transaction struct {
	TransactionID uint      `gorm:"primaryKey"`
	UserID        uint      `gorm:"not null;index"` // Kolom biasa tanpa foreign key
	Amount        int       `gorm:"not null"`
	Type          string    `gorm:"not null"`
	Category      string    `gorm:"not null"`
	Descriptions  string    `gorm:"not null"`
	Date          time.Time `gorm:"not null"`
}
