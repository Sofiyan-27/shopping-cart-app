package models

import "time"

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
