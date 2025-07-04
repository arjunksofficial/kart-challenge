package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/arjunksofficial/kart-challenge/internal/config"
	"github.com/arjunksofficial/kart-challenge/internal/core/logger"
	"github.com/arjunksofficial/kart-challenge/internal/core/serror"
	"github.com/arjunksofficial/kart-challenge/internal/entities/orders/models"
	"github.com/arjunksofficial/kart-challenge/internal/entities/orders/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func TestHandler_CreateOrder(t *testing.T) {
	config.SetConfig(&config.Config{
		Port: "8080",
	})
	t.Run("Success order creation", func(t *testing.T) {
		mockService := service.NewMockService(t)
		mockService.On("CreateOrder", mock.Anything, mock.Anything).Return(models.CreateOrderResponse{}, nil)
		handler := Handler{
			Service: mockService,
			logger: &logger.CustomLogger{
				Logger: zap.NewNop(),
			},
		}
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		requestBody := `{
			"couponCode": "BIRTHDAY",
			"items": [
				{
					"productId": "12345",
					"quantity": 2
				}
			]
		}`
		c.Request, _ = http.NewRequest(http.MethodPost, "/api/v1/orders", strings.NewReader(requestBody))

		handler.CreateOrder(c)

		assert.Equal(t, http.StatusOK, w.Code)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid request body", func(t *testing.T) {
		mockService := service.NewMockService(t)
		handler := Handler{
			Service: mockService,
			logger: &logger.CustomLogger{
				Logger: zap.NewNop(),
			},
		}
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		requestBody := `{"invalidField": "value"}`
		c.Request, _ = http.NewRequest(http.MethodPost, "/api/v1/orders", strings.NewReader(requestBody))

		handler.CreateOrder(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		mockService.AssertExpectations(t)
	})
	t.Run("Validation error", func(t *testing.T) {
		mockService := service.NewMockService(t)
		handler := Handler{
			Service: mockService,
			logger: &logger.CustomLogger{
				Logger: zap.NewNop(),
			},
		}
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		requestBody := `{
			"couponCode": "BIRTHDAY",
			"items": [
				{
					"productId": "12345",
					"quantity": -1 
				}
			]
		}`
		c.Request, _ = http.NewRequest(http.MethodPost, "/api/v1/orders", strings.NewReader(requestBody))
		handler.CreateOrder(c)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		mockService.AssertExpectations(t)
	})
	t.Run("Service unavailable", func(t *testing.T) {
		mockService := service.NewMockService(t)
		mockService.On("CreateOrder", mock.Anything, mock.Anything).Return(models.CreateOrderResponse{}, &serror.ServiceError{
			Code:  http.StatusBadRequest,
			Error: errors.New("bad coupon code"),
		})
		handler := Handler{
			Service: mockService,
			logger: &logger.CustomLogger{
				Logger: zap.NewNop(),
			},
		}
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		requestBody := `{
			"couponCode": "BIRTHDAY",
			"items": [
				{
					"productId": "12345",
					"quantity": 2
				}
			]
		}`
		c.Request, _ = http.NewRequest(http.MethodPost, "/api/v1/orders", strings.NewReader(requestBody))

		handler.CreateOrder(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"message":"bad coupon code"}`, w.Body.String())
		mockService.AssertExpectations(t)
	})
	t.Run("Error from service", func(t *testing.T) {
		mockService := service.NewMockService(t)
		mockService.On("CreateOrder", mock.Anything, mock.Anything).Return(models.CreateOrderResponse{}, &serror.ServiceError{
			Code:  http.StatusInternalServerError,
			Error: errors.New("some error"),
		})
		handler := Handler{
			Service: mockService,
			logger: &logger.CustomLogger{
				Logger: zap.NewNop(),
			},
		}
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		requestBody := `{
			"couponCode": "BIRTHDAY",
			"items": [
				{
					"productId": "12345",
					"quantity": 2
				}
			]
		}`
		c.Request, _ = http.NewRequest(http.MethodPost, "/api/v1/orders", strings.NewReader(requestBody))

		handler.CreateOrder(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		mockService.AssertExpectations(t)
	})
}
