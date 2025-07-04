package models

import (
	"time"

	"github.com/arjunksofficial/kart-challenge/internal/entities/products/models"
)

type Order struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	CouponCode string    `gorm:"type:varchar(10)" json:"coupon_code"`
	CreatedAt  time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp" json:"updated_at"`
}

type OrderItem struct {
	ID        int            `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID   int            `gorm:"not null" json:"order_id"`
	Order     Order          `gorm:"foreignKey:OrderID" json:"order"`
	ProductID string         `gorm:"not null" json:"product_id"`
	Product   models.Product `gorm:"foreignKey:ProductID" json:"product"`
	Quantity  int            `gorm:"not null" json:"quantity"`
	CreatedAt time.Time      `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp" json:"updated_at"`
}
