package store

import (
	"github.com/arjunksofficial/kart-challenge/internal/database"
	"github.com/arjunksofficial/kart-challenge/internal/entities/products/models"
	"gorm.io/gorm"
)

type Store interface {
	ListProducts() ([]models.Product, error)
}

type store struct {
	db *gorm.DB
}

func GetStore() Store {
	return &store{
		db: database.GetPostgresDB(),
	}
}
func (s *store) ListProducts() ([]models.Product, error) {
	var products []models.Product
	if err := s.db.Preload("Images").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
