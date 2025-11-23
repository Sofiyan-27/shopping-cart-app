package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"shopping-cart-app/backend/config"
	"shopping-cart-app/backend/models"
)

// POST /seed/items
func SeedItems(c *gin.Context) {
	items := []models.Item{
		{Name: "Soap", Status: "AVAILABLE"},
		{Name: "Shampoo", Status: "AVAILABLE"},
		{Name: "Toothpaste", Status: "AVAILABLE"},
		{Name: "Juice", Status: "AVAILABLE"},
		{Name: "Biscuits", Status: "AVAILABLE"},
	}

	for _, item := range items {
		config.DB.Create(&item)
	}

	c.JSON(http.StatusOK, gin.H{"message": "items seeded"})
}
