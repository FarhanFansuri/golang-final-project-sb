package models

// type User struct {
// 	UserID    uint      `gorm:"primaryKey"`
// 	Username  string    `gorm:"type:varchar(255);not null"`
// 	Email     string    `gorm:"type:varchar(255);not null"`
// 	Password  string    `gorm:"type:varchar(255);not null"`
// 	CreatedAt time.Time `gorm:"default:current_timestamp"`
// }

// type Transaction struct {
// 	TransactionID uint      `gorm:"primaryKey"`
// 	UserID        uint      `gorm:"not null;index"` // Kolom biasa tanpa foreign key
// 	Amount        int       `gorm:"not null"`
// 	Type          string    `gorm:"not null"`
// 	Category      string    `gorm:"not null"`
// 	Descriptions  string    `gorm:"not null"`
// 	Date          time.Time `gorm:"not null"`
// }

type User struct {
	UserID   uint   `gorm:"primaryKey"`
	Username string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null;unique"` // Email dibuat unik
	Password string `gorm:"type:varchar(255);not null"`
	// Bisa NULL
}

// type Transaction struct {
// 	TransactionID uint    `gorm:"primaryKey"`
// 	UserID        uint    `gorm:"not null;index"` // Kolom biasa tanpa foreign key
// 	Amount        int     `gorm:"not null"`
// 	Type          string  `gorm:"type:varchar(100);not null"`
// 	Category      string  `gorm:"type:varchar(100);not null"`
// 	Descriptions  string  `gorm:"type:varchar(255);not null"`
// 	Date          []uint8 `gorm:"not null"`
// }

type Transaction struct {
	ID           uint    `json:"id"`
	UserID       uint    `json:"UserID"`
	Amount       float64 `json:"Amount"`
	Type         string  `json:"Type"`
	Category     string  `json:"Category"`
	Descriptions string  `json:"Descriptions"`
	Date         string  `json:"Date"`
}
