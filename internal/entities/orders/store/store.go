package store

import (
	"context"

	"github.com/arjunksofficial/kart-challenge/internal/database"
	"github.com/arjunksofficial/kart-challenge/internal/entities/orders/models"
	"gorm.io/gorm"
)

type Store interface {
	CreateOrder(ctx context.Context, order *models.Order) error
	GetOrderByID(ctx context.Context, id int) (*models.Order, error)

	CreateOrderItems(ctx context.Context, orderItems []models.OrderItem) error
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

func (s *store) CreateOrder(ctx context.Context, order *models.Order) error {
	if err := s.db.WithContext(ctx).Create(order).Error; err != nil {
		return err
	}
	return nil
}

func (s *store) GetOrderByID(ctx context.Context, id int) (*models.Order, error) {
	var order models.Order
	if err := s.db.WithContext(ctx).First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (s *store) CreateOrderItems(ctx context.Context, orderItems []models.OrderItem) error {
	if len(orderItems) == 0 {
		return nil // No order items to create
	}

	// Use the Create method to insert multiple records
	result := s.db.WithContext(ctx).Create(&orderItems)
	if result.Error != nil {
		return result.Error // Return any error that occurred during the creation
	}

	return nil // Return nil if the operation was successful
}
