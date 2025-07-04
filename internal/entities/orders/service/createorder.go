package service

import (
	"context"
	"net/http"

	"github.com/arjunksofficial/kart-challenge/internal/core/serror"
	"github.com/arjunksofficial/kart-challenge/internal/entities/orders/models"
	"github.com/pkg/errors"
)

func (s *service) CreateOrder(ctx context.Context, req models.CreateOrderRequest) (models.Order, *serror.ServiceError) {
	order := models.Order{
		CouponCode: req.CouponCode,
	}

	if err := s.db.CreateOrder(&order); err != nil {
		return models.Order{}, &serror.ServiceError{
			Code:  http.StatusInternalServerError,
			Error: errors.Wrap(err, "failed to create order"),
		}

	}
	orderItems := []models.OrderItem{}
	for _, item := range req.Items {
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
		orderItems = append(orderItems, orderItem)
	}

	if err := s.db.CreateOrderItems(orderItems); err != nil {
		return models.Order{}, &serror.ServiceError{
			Code:  http.StatusInternalServerError,
			Error: errors.Wrap(err, "failed to create order item"),
		}
	}
	return order, nil
}
