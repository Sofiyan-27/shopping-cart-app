package models

import "time"

type Order struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserID    uint      `json:"user_id"`
	CartID    uint      `json:"cart_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
