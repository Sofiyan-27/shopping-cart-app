package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"shopping-cart-app/backend/config"
	"shopping-cart-app/backend/models"
)

func GetItems(c *gin.Context) {
	var items []models.Item
	if err := config.DB.Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch items"})
		return
	}

	c.JSON(http.StatusOK, items)
}
