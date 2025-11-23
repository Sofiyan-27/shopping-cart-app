package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"shopping-cart-app/backend/config"
	"shopping-cart-app/backend/models"
)

// POST /orders  (checkout)
func CreateOrder(c *gin.Context) {
	userIdVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found in context"})
		return
	}
	userID, ok := userIdVal.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id type"})
		return
	}

	// find active cart
	var cart models.Cart
	if err := config.DB.Where("user_id = ? AND status = ?", userID, "ACTIVE").First(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no active cart"})
		return
	}

	// ensure cart has items
	var count int64
	config.DB.Model(&models.CartItem{}).Where("cart_id = ?", cart.ID).Count(&count)
	if count == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cart is empty"})
		return
	}

	// create order
	order := models.Order{
		UserID: userID,
		CartID: cart.ID,
	}
	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
		return
	}

	// mark cart as checked out
	cart.Status = "CHECKED_OUT"
	config.DB.Save(&cart)

	c.JSON(http.StatusOK, gin.H{
		"message":  "order created",
		"order_id": order.ID,
	})
}

// GET /orders  (order history)
func GetOrders(c *gin.Context) {
	userIdVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found in context"})
		return
	}
	userID, ok := userIdVal.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id type"})
		return
	}

	var orders []models.Order
	if err := config.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
