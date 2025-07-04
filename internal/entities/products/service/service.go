package service

import (
	"context"

	"github.com/arjunksofficial/kart-challenge/internal/core/logger"
	"github.com/arjunksofficial/kart-challenge/internal/core/serror"
	"github.com/arjunksofficial/kart-challenge/internal/entities/products/models"
	"github.com/arjunksofficial/kart-challenge/internal/entities/products/store"
)

type service struct {
	db     store.Store
	logger *logger.CustomLogger
}

type Service interface {
	ListProducts(ctx context.Context) ([]models.ProductResponse, *serror.ServiceError)
	GetProductByID(ctx context.Context, id string) (models.ProductResponse, *serror.ServiceError)
}

func GetService() Service {
	return &service{
		db:     store.Get(),
		logger: logger.GetLogger(),
	}
}
