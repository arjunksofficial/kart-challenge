package service

import (
	"context"

	"github.com/arjunksofficial/kart-challenge/internal/core/serror"
	"github.com/arjunksofficial/kart-challenge/internal/entities/orders/models"
	"github.com/arjunksofficial/kart-challenge/internal/entities/orders/store"
	productstore "github.com/arjunksofficial/kart-challenge/internal/entities/products/store"
	promocodestore "github.com/arjunksofficial/kart-challenge/internal/entities/promocode/store"
)

type service struct {
	store          store.Store
	productstore   productstore.Store
	promocodestore promocodestore.Cache
}

type Service interface {
	CreateOrder(ctx context.Context, req models.CreateOrderRequest) (models.CreateOrderResponse, *serror.ServiceError)
}

var svc Service

func New() Service {
	return &service{
		store:          store.Get(),
		productstore:   productstore.Get(),
		promocodestore: promocodestore.Get(),
	}
}
func Get() Service {
	if svc == nil {
		svc = New()
	}
	return svc
}
