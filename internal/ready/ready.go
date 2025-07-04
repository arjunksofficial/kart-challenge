package ready

import (
	"net/http"

	"github.com/arjunksofficial/kart-challenge/internal/config"
	"github.com/arjunksofficial/kart-challenge/internal/rediscli"
	"github.com/gin-gonic/gin"
)

func ReadyHandler(w http.ResponseWriter, r *http.Request) {
	// Respond with a simple message indicating readiness
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Ready"))
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
	// check for redis connection or any other service readiness checks here
	cfg := config.GetConfig()
	if !cfg.IsReady() {
		http.Error(w, "Service not ready", http.StatusServiceUnavailable)
		return
	}
	err = rediscli.GetRedisClient().Ping(r.Context()).Err() // Check Redis connection
	if err != nil {
		http.Error(w, "Redis not ready: "+err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.Write([]byte("Service is ready"))
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK) // Ensure the status code is set to 200 OK
}

func Ready(c *gin.Context) {
	cfg := config.GetConfig()
	if !cfg.IsReady() {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Service not ready"})
		return
	}
	err := rediscli.GetRedisClient().Ping(c.Request.Context()).Err() // Check Redis connection
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Redis not ready: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Service is ready"})
}
