package models

import "time"

type Product struct {
	ProductMeta
	Images    ProductImages `gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE" json:"image"`
	CreatedAt time.Time     `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time     `gorm:"type:timestamp" json:"updated_at"`
}

type ProductMeta struct {
	ID       string  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string  `gorm:"type:varchar(255);not null" json:"name"`
	Category string  `gorm:"type:varchar(100);not null" json:"category"`
	Price    float64 `gorm:"type:decimal(10,2);not null" json:"price"`
}

type ProductMetaWithImages struct {
	ProductMeta
	Images ProductImages `gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE" json:"images"`
}

type ProductImages struct {
	ProductID string `gorm:"primaryKey;type:text" json:"-"`
	Thumbnail string `gorm:"type:text" json:"thumbnail"`
	Mobile    string `gorm:"type:text" json:"mobile"`
	Tablet    string `gorm:"type:text" json:"tablet"`
	Desktop   string `gorm:"type:text" json:"desktop"`
}
