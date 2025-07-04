package ready

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arjunksofficial/kart-challenge/internal/config"
	"github.com/arjunksofficial/kart-challenge/internal/rediscli"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
)

func TestReady(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Service is ready", func(t *testing.T) {
		config.SetConfig(&config.Config{
			Port: "8080",
		})
		mockRedisClient, mock := redismock.NewClientMock()
		mock.ExpectPing().SetVal("PONG") // Mock Redis ping response
		rediscli.SetRedisClient(mockRedisClient)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodGet, "/ready", nil)

		Ready(c)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"message":"Service is ready"}`, w.Body.String())
		mock.ExpectationsWereMet()
	})
	t.Run("Service not ready", func(t *testing.T) {

		config.SetConfig(&config.Config{
			Port: "",
		}) // Simulate service not ready by not setting the port
		mockRedisClient, mock := redismock.NewClientMock()
		rediscli.SetRedisClient(mockRedisClient)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodGet, "/ready", nil)

		Ready(c)

		assert.Equal(t, http.StatusServiceUnavailable, w.Code)
		assert.JSONEq(t, `{"error":"Service not ready"}`, w.Body.String())
		mock.ExpectationsWereMet()
	})
	t.Run("Redis not ready", func(t *testing.T) {
		config.SetConfig(&config.Config{
			Port: "8080",
		})
		mockRedisClient, mock := redismock.NewClientMock()
		mock.ExpectPing().SetErr(errors.New("some error"))
		rediscli.SetRedisClient(mockRedisClient)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodGet, "/ready", nil)

		Ready(c)

		assert.Equal(t, http.StatusServiceUnavailable, w.Code)
		assert.JSONEq(t, `{"error":"Redis not ready: some error"}`, w.Body.String())
		mock.ExpectationsWereMet()
	})
}
