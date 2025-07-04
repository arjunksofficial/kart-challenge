package main

import (
	"fmt"

	"github.com/arjunksofficial/kart-challenge/internal/core/logger"
	"github.com/arjunksofficial/kart-challenge/internal/database"
	ordermodels "github.com/arjunksofficial/kart-challenge/internal/entities/orders/models"
	productmodels "github.com/arjunksofficial/kart-challenge/internal/entities/products/models"
	"github.com/arjunksofficial/kart-challenge/internal/migrator"
)

func main() {
	logger := logger.GetLogger()
	logger.Info("Connecting to Postgres database...")
	db := database.GetPostgresDB()
	logger.Info("Connected to Postgres database successfully.")
	logger.Info("Starting migration...")
	// 🔧 Auto-migrate your models
	err := db.AutoMigrate(&productmodels.Product{}, &productmodels.ProductImages{}, &ordermodels.Order{}, &ordermodels.OrderItem{})
	if err != nil {
		panic("migration failed: " + err.Error())
	}
	// 🔧 Insert sample product
	err = db.Create(&migrator.SampleProducts).Error
	if err != nil {
		panic("failed to insert sample products: " + err.Error())
	}
	fmt.Println("Migration completed.")
}
