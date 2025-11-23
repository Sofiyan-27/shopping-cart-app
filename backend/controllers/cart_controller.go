package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"shopping-cart-app/backend/config"
	"shopping-cart-app/backend/models"
)

type AddToCartRequest struct {
	ItemID uint `json:"item_id"`
}

func getOrCreateActiveCart(db *gorm.DB, userID uint) (*models.Cart, error) {
	var cart models.Cart
	err := db.Where("user_id = ? AND status = ?", userID, "ACTIVE").First(&cart).Error
	if err == nil {
		return &cart, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		newCart := models.Cart{
			UserID: userID,
			Name:   "Default Cart",
			Status: "ACTIVE",
		}
		if err := db.Create(&newCart).Error; err != nil {
			return nil, err
		}
		return &newCart, nil
	}
	return nil, err
}

func AddToCart(c *gin.Context) {
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

	var req AddToCartRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.ItemID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	var item models.Item
	if err := config.DB.First(&item, req.ItemID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "item not found"})
		return
	}

	cart, err := getOrCreateActiveCart(config.DB, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get or create cart"})
		return
	}

	cartItem := models.CartItem{
		CartID: cart.ID,
		ItemID: item.ID,
	}
	if err := config.DB.Create(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add to cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "item added to cart",
		"cart_id": cart.ID,
		"item_id": item.ID,
	})
}

func GetCart(c *gin.Context) {
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

	var cart models.Cart
	if err := config.DB.Where("user_id = ? AND status = ?", userID, "ACTIVE").First(&cart).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "no active cart"})
		return
	}

	var cartItems []models.CartItem
	if err := config.DB.Where("cart_id = ?", cart.ID).Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch cart items"})
		return
	}

	itemIDs := make([]uint, 0, len(cartItems))
	for _, ci := range cartItems {
		itemIDs = append(itemIDs, ci.ItemID)
	}

	c.JSON(http.StatusOK, gin.H{
		"cart_id":  cart.ID,
		"item_ids": itemIDs,
	})
}
