package store

import (
	"context"

	"github.com/arjunksofficial/kart-challenge/internal/database"
	"github.com/arjunksofficial/kart-challenge/internal/entities/products/models"
	"gorm.io/gorm"
)

type Store interface {
	ListProducts(ctx context.Context, filter models.ProductFilter) ([]models.Product, error)
	GetByID(ctx context.Context, id string) (*models.Product, error)
}

type store struct {
	db *gorm.DB
}

var postgresStore Store

func New() Store {
	return &store{
		db: database.GetPostgresDB(),
	}
}

func Get() Store {
	if postgresStore == nil {
		postgresStore = New()
	}
	return postgresStore
}

// ListProducts retrieves a list of products based on the provided filter.
func (s *store) ListProducts(ctx context.Context, filter models.ProductFilter) ([]models.Product, error) {
	var products []models.Product
	query := s.db.Preload("Images")
	if len(filter.ProductIDs) > 0 {
		query = query.Where("id IN ?", filter.ProductIDs)
	}
	if filter.Category != "" {
		query = query.Where("category = ?", filter.Category)
	}
	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return products, nil
}

func (s *store) GetByID(ctx context.Context, id string) (*models.Product, error) {
	var product models.Product
	if err := s.db.Preload("Images").First(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}
