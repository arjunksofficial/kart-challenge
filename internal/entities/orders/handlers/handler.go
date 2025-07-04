package handlers

import (
	"github.com/arjunksofficial/kart-challenge/internal/core/logger"
	"github.com/arjunksofficial/kart-challenge/internal/entities/orders/service"
)

type Handler struct {
	Service service.Service
	logger  *logger.CustomLogger
}

func New() *Handler {
	return &Handler{
		Service: service.GetService(),
		logger:  logger.GetLogger(),
	}
}
