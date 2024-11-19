package models

import "time"

// Struct untuk tabel menu
type Menu struct {
    MenuID  int    `json:"menu_id"`
    Name    string `json:"name"`
    Category string `json:"category"`
    Price   int    `json:"price"`
}

// Struct untuk tabel orders
type Orders struct {
    OrderID   int       `json:"order_id"`
    UserID    int       `json:"user_id"`
    OrderDate time.Time `json:"order_date"`
    Status    string    `json:"status"`
}

// Struct untuk tabel order_items
type OrderItems struct {
    OrderID    int `json:"order_id"`
    MenuID     int `json:"menu_id"`
    Quantity   int `json:"quantity"`
    TotalPrice int `json:"total_price"`
}

// Struct untuk tabel payments
type Payments struct {
    PaymentID     int       `json:"payment_id"`
    OrderID       int       `json:"order_id"`
    PaymentMethod string    `json:"payment_method"`
    PaymentDate   time.Time `json:"payment_date"`
    Status        string    `json:"status"`
}

// Struct untuk tabel users
type Users struct {
    UserID   int    `json:"user_id"`
    Username string `json:"username"`
    Password string `json:"password"`
    Role     string `json:"role"`
}
