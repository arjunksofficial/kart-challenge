package ready

import (
	"net/http"

	"github.com/arjunksofficial/kart-challenge/internal/config"
	"github.com/arjunksofficial/kart-challenge/internal/rediscli"
	"github.com/gin-gonic/gin"
)

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
