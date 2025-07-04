package api

import (
	"github.com/arjunksofficial/kart-challenge/internal/entities/orders"
	"github.com/arjunksofficial/kart-challenge/internal/entities/products"
	"github.com/arjunksofficial/kart-challenge/internal/ready"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(gin.ErrorLogger())
	router.GET("/", HealthCheck)
	// Health check endpoint
	router.GET("/health", HealthCheck)
	// Readiness check endpoint
	router.GET("/ready", ready.Ready)

	apiRoutes := router.Group("/api/v1")
	// Register entity-specific routes
	products.RegisterRoutes(apiRoutes)
	orders.RegisterRoutes(apiRoutes)

	return router
}

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}
