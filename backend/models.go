package main

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CardID    int       `json:"card_id"`
	Role      int       `json:"role"` // Role can be 0=disabled, 1=admin, 2=guest
	LastLogin time.Time `json:"last_login"`
}

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description"`
}

type Inventory struct {
	ID          int       `json:"id"`
	BookID      int       `json:"book_id"`
	GuestID     int       `json:"guest_id"`
	Status      string    `json:"status"` // Status can be "in" or "out"
	LastUpdated time.Time `json:"last_updated"`
}
