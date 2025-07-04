package models

import "time"

type Product struct {
	ProductMeta
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}

type ProductMeta struct {
	ID       string  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string  `gorm:"type:varchar(255);not null" json:"name"`
	Category string  `gorm:"type:varchar(100);not null" json:"category"`
	Price    float64 `gorm:"type:decimal(10,2);not null" json:"price"`
}
