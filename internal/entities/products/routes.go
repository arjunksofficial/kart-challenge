package products

import (
	"github.com/arjunksofficial/kart-challenge/internal/entities/products/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(apiRoutes *gin.RouterGroup) {
	// Create a new h instance
	h := handlers.New()
	productRoutes := apiRoutes.Group("/products")
	{
		productRoutes.GET("", h.ListProducts)
		productRoutes.GET("/:id", h.GetProductByID)
	}
}
