package service

import (
	"context"
	"net/http"
	"testing"

	"github.com/arjunksofficial/kart-challenge/internal/config"
	"github.com/arjunksofficial/kart-challenge/internal/entities/orders/models"
	"github.com/arjunksofficial/kart-challenge/internal/entities/orders/store"
	productmodels "github.com/arjunksofficial/kart-challenge/internal/entities/products/models"
	productstore "github.com/arjunksofficial/kart-challenge/internal/entities/products/store"
	promocodestore "github.com/arjunksofficial/kart-challenge/internal/entities/promocode/store"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestService_CreateOrder(t *testing.T) {
	config.SetConfig(&config.Config{
		Port: "8080",
	})
	t.Run("Successful order creation", func(t *testing.T) {
		mockstore := store.NewMockStore(t)
		mockstore.On("CreateOrder", mock.Anything, mock.Anything).
			Run(func(args mock.Arguments) {
				order := args.Get(1).(*models.Order)
				order.ID = 123
			}).Return(nil)
		mockstore.On("CreateOrderItems", mock.Anything, mock.Anything).Return(nil)
		mockpromocodecache := promocodestore.NewMockCache(t)
		mockpromocodecache.On("IsPresentInSet", mock.Anything, "valid_promo_codes", "BIRTHDAY").Return(true, nil)
		mockproductstore := productstore.NewMockStore(t)
		mockproductstore.On("ListProducts", mock.Anything, mock.Anything).Return([]productmodels.Product{
			{
				ProductMeta: productmodels.ProductMeta{
					ID:       "12345",
					Name:     "Test Product",
					Category: "Test Category",
					Price:    100.0,
				},
				Images: productmodels.ProductImages{
					ProductID: "12345",
					Thumbnail: "http://example.com/thumbnail.jpg",
					Mobile:    "http://example.com/mobile.jpg",
					Tablet:    "http://example.com/tablet.jpg",
					Desktop:   "http://example.com/desktop.jpg",
				},
			},
		}, nil)

		service := &service{
			store:          mockstore,
			productstore:   mockproductstore,
			promocodestore: mockpromocodecache,
		}
		req := models.CreateOrderRequest{
			CouponCode: "BIRTHDAY",
			Items: []models.Item{
				{
					ProductID: "12345",
					Quantity:  2,
				},
			},
		}
		resp, sErr := service.CreateOrder(context.Background(), req)
		assert.Nil(t, sErr)
		assert.Equal(t, models.CreateOrderResponse{
			CouponCode: "BIRTHDAY",
			ID:         123,
			Items: []models.Item{
				{
					ProductID: "12345",
					Quantity:  2,
				},
			},
			Products: []productmodels.ProductMetaWithImages{
				{
					ProductMeta: productmodels.ProductMeta{
						ID:       "12345",
						Name:     "Test Product",
						Category: "Test Category",
						Price:    100.0,
					},
					Images: productmodels.ProductImages{
						ProductID: "12345",
						Thumbnail: "http://example.com/thumbnail.jpg",
						Mobile:    "http://example.com/mobile.jpg",
						Tablet:    "http://example.com/tablet.jpg",
						Desktop:   "http://example.com/desktop.jpg",
					},
				},
			},
		}, resp)

		mockstore.AssertExpectations(t)
		mockpromocodecache.AssertExpectations(t)
		mockproductstore.AssertExpectations(t)

	})

	t.Run("Order creation with invalid coupon code", func(t *testing.T) {
		mockstore := store.NewMockStore(t)
		mockpromocodecache := promocodestore.NewMockCache(t)
		mockpromocodecache.On("IsPresentInSet", mock.Anything, "valid_promo_codes", "INVALID").Return(false, nil)

		service := &service{
			store:          mockstore,
			promocodestore: mockpromocodecache,
		}
		req := models.CreateOrderRequest{
			CouponCode: "INVALID",
			Items: []models.Item{
				{
					ProductID: "12345",
					Quantity:  2,
				},
			},
		}
		resp, sErr := service.CreateOrder(context.Background(), req)
		assert.NotNil(t, sErr)
		assert.Equal(t, http.StatusBadRequest, sErr.Code)
		assert.Equal(t, "invalid coupon code", sErr.Error.Error())
		assert.Equal(t, models.CreateOrderResponse{}, resp)

		mockstore.AssertExpectations(t)
		mockpromocodecache.AssertExpectations(t)
	})
	t.Run("Order creation failed due to promocode cache error", func(t *testing.T) {
		mockstore := store.NewMockStore(t)
		mockpromocodecache := promocodestore.NewMockCache(t)
		mockpromocodecache.On("IsPresentInSet", mock.Anything, "valid_promo_codes", "BIRTHDAY").Return(false, assert.AnError)

		service := &service{
			store:          mockstore,
			promocodestore: mockpromocodecache,
		}
		req := models.CreateOrderRequest{
			CouponCode: "BIRTHDAY",
			Items: []models.Item{
				{
					ProductID: "12345",
					Quantity:  2,
				},
			},
		}
		resp, sErr := service.CreateOrder(context.Background(), req)
		assert.NotNil(t, sErr)
		assert.Equal(t, "failed to check coupon code validity: "+assert.AnError.Error(), sErr.Error.Error())
		assert.Equal(t, models.CreateOrderResponse{}, resp)

		mockstore.AssertExpectations(t)
		mockpromocodecache.AssertExpectations(t)
	})
	t.Run("Order creation failed due to store error", func(t *testing.T) {
		mockstore := store.NewMockStore(t)
		mockstore.On("CreateOrder", mock.Anything, mock.Anything).Return(assert.AnError)
		mockpromocodecache := promocodestore.NewMockCache(t)
		mockpromocodecache.On("IsPresentInSet", mock.Anything, "valid_promo_codes", "BIRTHDAY").Return(true, nil)

		service := &service{
			store:          mockstore,
			promocodestore: mockpromocodecache,
		}
		req := models.CreateOrderRequest{
			CouponCode: "BIRTHDAY",
			Items: []models.Item{
				{
					ProductID: "12345",
					Quantity:  2,
				},
			},
		}
		resp, sErr := service.CreateOrder(context.Background(), req)
		assert.NotNil(t, sErr)
		assert.Equal(t, "failed to create order: "+assert.AnError.Error(), sErr.Error.Error())
		assert.Equal(t, models.CreateOrderResponse{}, resp)

		mockstore.AssertExpectations(t)
		mockpromocodecache.AssertExpectations(t)
	})
	t.Run("Order creation failed due to order item creation error", func(t *testing.T) {
		mockstore := store.NewMockStore(t)
		mockstore.On("CreateOrder", mock.Anything, mock.Anything).
			Run(func(args mock.Arguments) {
				order := args.Get(1).(*models.Order)
				order.ID = 123
			}).Return(nil)
		mockstore.On("CreateOrderItems", mock.Anything, mock.Anything).Return(assert.AnError)
		mockpromocodecache := promocodestore.NewMockCache(t)
		mockpromocodecache.On("IsPresentInSet", mock.Anything, "valid_promo_codes", "BIRTHDAY").Return(true, nil)

		service := &service{
			store:          mockstore,
			promocodestore: mockpromocodecache,
		}
		req := models.CreateOrderRequest{
			CouponCode: "BIRTHDAY",
			Items: []models.Item{
				{
					ProductID: "12345",
					Quantity:  2,
				},
			},
		}
		resp, sErr := service.CreateOrder(context.Background(), req)
		assert.NotNil(t, sErr)
		assert.Equal(t, "failed to create order item: "+assert.AnError.Error(), sErr.Error.Error())
		assert.Equal(t, models.CreateOrderResponse{}, resp)

		mockstore.AssertExpectations(t)
		mockpromocodecache.AssertExpectations(t)
	})
	t.Run("Order creation failed due to product store error", func(t *testing.T) {
		mockstore := store.NewMockStore(t)
		mockstore.On("CreateOrder", mock.Anything, mock.Anything).
			Run(func(args mock.Arguments) {
				order := args.Get(1).(*models.Order)
				order.ID = 123
			}).Return(nil)
		mockstore.On("CreateOrderItems", mock.Anything, mock.Anything).Return(nil)
		mockpromocodecache := promocodestore.NewMockCache(t)
		mockpromocodecache.On("IsPresentInSet", mock.Anything, "valid_promo_codes", "BIRTHDAY").Return(true, nil)
		mockproductstore := productstore.NewMockStore(t)
		mockproductstore.On("ListProducts", mock.Anything, mock.Anything).Return(nil, assert.AnError)

		service := &service{
			store:          mockstore,
			productstore:   mockproductstore,
			promocodestore: mockpromocodecache,
		}
		req := models.CreateOrderRequest{
			CouponCode: "BIRTHDAY",
			Items: []models.Item{
				{
					ProductID: "12345",
					Quantity:  2,
				},
			},
		}
		resp, sErr := service.CreateOrder(context.Background(), req)
		assert.NotNil(t, sErr)
		assert.Equal(t, "failed to list products: "+assert.AnError.Error(), sErr.Error.Error())
		assert.Equal(t, models.CreateOrderResponse{}, resp)

		mockstore.AssertExpectations(t)
		mockpromocodecache.AssertExpectations(t)
		mockproductstore.AssertExpectations(t)
	})
}
