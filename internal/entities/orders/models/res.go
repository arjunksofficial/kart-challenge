package models

import (
	"github.com/arjunksofficial/kart-challenge/internal/entities/products/models"
)

type CreateOrderResponse struct {
	ID       int                  `json:"id"`
	Items    []Item               `json:"items"`
	Products []models.ProductMeta `json:"products"`
}
