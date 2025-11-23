package models

import "time"

type CartItem struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CartID    uint      `json:"cart_id"`
	ItemID    uint      `json:"item_id"`
	CreatedAt time.Time `json:"created_at"`
}
