package api

import (
	"github.com/arjunksofficial/kart-challenge/internal/entities/orders"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(gin.ErrorLogger())
	apiRoutes := router.Group("/api/v1")
	// Define your routes here
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	orders.RegisterRoutes(apiRoutes)
	return router
}
