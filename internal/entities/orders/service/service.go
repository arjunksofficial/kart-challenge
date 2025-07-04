package service

import (
	"context"

	"github.com/arjunksofficial/kart-challenge/internal/core/logger"
	"github.com/arjunksofficial/kart-challenge/internal/core/serror"
	"github.com/arjunksofficial/kart-challenge/internal/entities/orders/models"
	"github.com/arjunksofficial/kart-challenge/internal/entities/orders/store"
)

type service struct {
	db     store.Store
	logger *logger.CustomLogger
}

type Service interface {
	CreateOrder(ctx context.Context, req models.CreateOrderRequest) (models.Order, *serror.ServiceError)
}

func GetService() Service {
	return &service{
		db:     store.GetStore(),
		logger: logger.GetLogger(),
	}
}
