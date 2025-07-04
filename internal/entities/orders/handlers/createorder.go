package handlers

import (
	"net/http"

	"github.com/arjunksofficial/kart-challenge/internal/core/responsehelper"
	"github.com/arjunksofficial/kart-challenge/internal/entities/orders/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//	CreateOrder creates a new order
//
// @Summary		Create a new order
// @Description	Create a new order
// @ID			create-order
// @Tags		orders
// @Accept		json
// @Produce		json
// @Param		start_date	 query		string                   	  false	"Start Date"
// @Param		end_date	 query		string                   	  false	"End Date"
// @Success		200	         {object}	models.CreateOrderResp
// @Failure		400	         {object}	responsehelper.CommonResponse
// @Failure		500	         {object}	responsehelper.CommonResponse
// @Router		/api/v1/orders [post]
func (h *Handler) CreateOrder(c *gin.Context) {
	ctx := c.Request.Context()
	var req models.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Failed to bind request", zap.Error(err))
		c.JSON(http.StatusBadRequest, responsehelper.NewCommonResponse("Invalid request body"))
		return
	}
	// Validate the request
	if err := req.Validate(); err != nil {
		h.logger.Error("Invalid request", zap.Error(err))
		c.JSON(http.StatusBadRequest, responsehelper.NewCommonResponse(err.Error()))
		return
	}
	// Create the order
	resp, sErr := h.Service.CreateOrder(ctx, req)
	if sErr != nil {
		if sErr.Code >= http.StatusInternalServerError {
			h.logger.Error("Error fetching revenue by category", zap.Error(sErr.Error))
			c.JSON(sErr.Code, responsehelper.NewCommonResponse("Internal Server Error"))
			return
		}
		c.JSON(sErr.Code, responsehelper.NewCommonResponse(sErr.Error.Error()))
		return
	}
	c.JSON(http.StatusOK, resp)
}
