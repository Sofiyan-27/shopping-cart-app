package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"shopping-cart-app/backend/config"
	"shopping-cart-app/backend/routes"
)

func main() {
	// connect to database
	config.ConnectDB()
	defer config.DB.Close()

	// create gin router
	r := gin.Default()

	// enable CORS so React (different port) can call this API
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{
		"Origin",
		"Content-Type",
		"Accept",
		"Authorization",
	}

	r.Use(cors.New(corsConfig))
	// register all routes (we'll define this next)
	routes.RegisterRoutes(r)

	// start server on :8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
