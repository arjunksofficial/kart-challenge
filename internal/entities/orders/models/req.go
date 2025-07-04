package models

import "errors"

type CreateOrderRequest struct {
	CouponCode string `json:"couponCode"`
	Items      []Item `json:"items" binding:"required"`
}

type Item struct {
	ProductID string `json:"productId" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
}

var (
	ErrItemsRequired     = errors.New("items are required")
	ErrInvalidQuantity   = errors.New("quantity must be greater than zero")
	ErrProductIDRequired = errors.New("product ID is required")
)

func (r *CreateOrderRequest) Validate() error {
	if len(r.Items) == 0 {
		return ErrItemsRequired
	}
	for _, item := range r.Items {
		if item.Quantity <= 0 {
			return ErrInvalidQuantity
		}
		if item.ProductID == "" {
			return ErrProductIDRequired
		}
	}
	return nil
}
