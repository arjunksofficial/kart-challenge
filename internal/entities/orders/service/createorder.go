package service

import (
	"context"
	"net/http"

	"github.com/arjunksofficial/kart-challenge/internal/core/serror"
	"github.com/arjunksofficial/kart-challenge/internal/entities/orders/models"
	productmodels "github.com/arjunksofficial/kart-challenge/internal/entities/products/models"
	"github.com/pkg/errors"
)

func (s *service) CreateOrder(
	ctx context.Context, req models.CreateOrderRequest,
) (models.CreateOrderResponse, *serror.ServiceError) {
	order := models.Order{
		CouponCode: req.CouponCode,
	}
	if req.CouponCode != "" {
		// Check if the coupon code is valid
		isValid, err := s.promocodestore.IsPresentInSet(ctx, "valid_promo_codes", req.CouponCode)
		if err != nil {
			return models.CreateOrderResponse{}, &serror.ServiceError{
				Code:  http.StatusInternalServerError,
				Error: errors.Wrap(err, "failed to check coupon code validity"),
			}
		}
		if !isValid {
			return models.CreateOrderResponse{}, &serror.ServiceError{
				Code:  http.StatusBadRequest,
				Error: errors.New("invalid coupon code"),
			}
		}
	}
	if err := s.store.CreateOrder(ctx, &order); err != nil {
		return models.CreateOrderResponse{}, &serror.ServiceError{
			Code:  http.StatusInternalServerError,
			Error: errors.Wrap(err, "failed to create order"),
		}
	}
	resp := models.CreateOrderResponse{
		CouponCode: req.CouponCode,
		ID:         order.ID,
	}
	productIDs := []string{}
	orderItems := []models.OrderItem{}
	for _, item := range req.Items {
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
		orderItems = append(orderItems, orderItem)
		resp.Items = append(resp.Items, models.Item{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
		productIDs = append(productIDs, item.ProductID)
	}
	if err := s.store.CreateOrderItems(ctx, orderItems); err != nil {
		return models.CreateOrderResponse{}, &serror.ServiceError{
			Code:  http.StatusInternalServerError,
			Error: errors.Wrap(err, "failed to create order item"),
		}
	}
	products, err := s.productstore.ListProducts(ctx, productmodels.ProductFilter{
		ProductIDs: productIDs,
	})

	if err != nil {
		return models.CreateOrderResponse{}, &serror.ServiceError{
			Code:  http.StatusInternalServerError,
			Error: errors.Wrap(err, "failed to list products"),
		}
	}
	if len(products) == 0 {
		return models.CreateOrderResponse{}, &serror.ServiceError{
			Code:  http.StatusNotFound,
			Error: errors.New("no products found for the given IDs"),
		}
	}
	for _, product := range products {
		resp.Products = append(resp.Products, productmodels.ProductMetaWithImages{
			ProductMeta: product.ProductMeta,
			Images:      product.Images,
		})
	}
	return resp, nil
}
