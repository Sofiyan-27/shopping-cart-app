package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"shopping-cart-app/backend/models"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open("sqlite3", "app.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(
		&models.User{},
		&models.Item{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
	)

	DB = db
}
