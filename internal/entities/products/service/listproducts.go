package service

import (
	"context"
	"net/http"

	"github.com/arjunksofficial/kart-challenge/internal/core/serror"
	"github.com/arjunksofficial/kart-challenge/internal/entities/products/models"
	"github.com/pkg/errors"
)

func (s *service) ListProducts(ctx context.Context) ([]models.ProductResponse, *serror.ServiceError) {
	products, err := s.db.ListProducts()
	if err != nil {
		return nil, &serror.ServiceError{
			Code:  http.StatusInternalServerError,
			Error: errors.Wrap(err, "failed to list products"),
		}
	}

	return models.MapProductsToResponse(products), nil
}
