// cmd/promoimporter/main.go
package main

import (
	"log"
	"time"

	"github.com/arjunksofficial/kart-challenge/internal/config"
	"github.com/arjunksofficial/kart-challenge/internal/promoimporter"
)

func main() {
	log.Println("🚀 Starting Promo Importer...")
	startTime := time.Now()
	config.LoadPromoImporterConfig()
	promoimporter := promoimporter.New()
	if err := promoimporter.Run(); err != nil {
		log.Fatalf("❌ Promo Importer failed: %v", err)
	}
	log.Printf("✅ Promo Importer completed in %v", time.Since(startTime))
}
