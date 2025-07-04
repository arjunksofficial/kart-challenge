package orders

import (
	"github.com/arjunksofficial/kart-challenge/internal/entities/orders/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(apiRoutes *gin.RouterGroup) {
	// Create a new h instance
	h := handlers.New()
	orderRoutes := apiRoutes.Group("/orders")
	{
		orderRoutes.POST("", h.CreateOrder)
	}
}
