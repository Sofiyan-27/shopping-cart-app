package routes

import (
	"github.com/gin-gonic/gin"

	"shopping-cart-app/backend/controllers"
	"shopping-cart-app/backend/middleware"
)

func RegisterRoutes(r *gin.Engine) {

	r.POST("/users", controllers.Signup)
	r.POST("/users/login", controllers.Login)
	r.POST("/seed/items", controllers.SeedItems)

	auth := r.Group("/")
	auth.Use(middleware.Auth())

	auth.GET("/items", controllers.GetItems)
	auth.POST("/carts", controllers.AddToCart)
	auth.GET("/carts", controllers.GetCart)
	auth.POST("/orders", controllers.CreateOrder)
	auth.GET("/orders", controllers.GetOrders)
}
