package store

import (
	"github.com/arjunksofficial/kart-challenge/internal/database"
	"github.com/arjunksofficial/kart-challenge/internal/entities/products/models"
	"gorm.io/gorm"
)

type Store interface {
	BulkCreate(customers []models.Product) error
	BulkCreateOrUpdate(customers []models.Product) error
}

type store struct {
	db *gorm.DB
}

func GetStore() Store {
	return &store{
		db: database.GetPostgresDB(),
	}
}

func (s *store) BulkCreate(customers []models.Product) error {
	if len(customers) == 0 {
		return nil // No customers to create
	}

	// Use the Create method to insert multiple records
	result := s.db.Create(&customers)
	if result.Error != nil {
		return result.Error // Return any error that occurred during the creation
	}

	return nil // Return nil if the operation was successful
}

func (s *store) BulkCreateOrUpdate(customers []models.Product) error {
	if len(customers) == 0 {
		return nil // No customers to create or update
	}

	// Use the Save method to insert or update multiple records
	result := s.db.Save(&customers)
	if result.Error != nil {
		return result.Error // Return any error that occurred during the save operation
	}

	return nil // Return nil if the operation was successful
}
