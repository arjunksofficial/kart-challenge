package models

import (
	"github.com/arjunksofficial/kart-challenge/internal/entities/products/models"
)

type CreateOrderResponse struct {
	ID         int                            `json:"id"`
	CouponCode string                         `json:"couponCode"`
	Items      []Item                         `json:"items"`
	Products   []models.ProductMetaWithImages `json:"products"`
}
