//	@title			Lumel Assignment
//	@version		1.0
//	@description	This is a sample server for serving revenue stats

//	@contact.name	API Support

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8000
//	@BasePath	/api/v1

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/arjunksofficial/kart-challenge/internal/api"
	"github.com/arjunksofficial/kart-challenge/internal/config"
)

func main() {
	log.Println("Starting API server...")
	go func() {
		log.Println("Starting pprof server on :6060")
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatalf("Failed to start pprof server: %v", err)
		}
	}()

	config := config.GetConfig()

	log.Printf("Configuration loaded: %+v", config)
	r := api.GetRouter()
	if err := r.Run(":" + config.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
