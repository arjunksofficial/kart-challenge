package service

import (
	"context"
	"net/http"

	"github.com/arjunksofficial/kart-challenge/internal/core/serror"
	"github.com/arjunksofficial/kart-challenge/internal/entities/products/models"
	"github.com/pkg/errors"
)

func (s *service) GetProductByID(ctx context.Context, id string) (models.ProductResponse, *serror.ServiceError) {
	product, err := s.db.GetByID(ctx, id)
	if err != nil {
		return models.ProductResponse{}, &serror.ServiceError{
			Code:  http.StatusInternalServerError,
			Error: errors.Wrap(err, "failed to get product by ID"),
		}
	}
	if product == nil {
		return models.ProductResponse{}, &serror.ServiceError{
			Code:  http.StatusNotFound,
			Error: errors.New("product not found"),
		}
	}
	return models.MapProductToResponse(*product), nil
}
