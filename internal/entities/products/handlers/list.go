package handlers

import (
	"net/http"

	"github.com/arjunksofficial/kart-challenge/internal/core/responsehelper"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ListProducts lists all products
//
// @Summary		List all products
// @Description	List all products
// @ID			list-products
// @Tags		products
// @Accept		json
// @Produce		json
// @Param		start_date	 query		string                   	  false	"Start Date"
// @Param		end_date	 query		string                   	  false	"End Date"
// @Success		200	         {object}	models.CreateOrderResp
// @Failure		400	         {object}	responsehelper.CommonResponse
// @Failure		500	         {object}	responsehelper.CommonResponse
// @Router		/api/v1/orders [post]
func (h *Handler) ListProducts(c *gin.Context) {
	ctx := c.Request.Context()
	resp, sErr := h.Service.ListProducts(ctx)
	if sErr != nil {
		if sErr.Code >= http.StatusInternalServerError {
			h.logger.Error("Error fetching products", zap.Error(sErr.Error))
			c.JSON(sErr.Code, responsehelper.NewCommonResponse("Internal Server Error"))
			return
		}
		c.JSON(sErr.Code, responsehelper.NewCommonResponse(sErr.Error.Error()))
		return
	}
	c.JSON(http.StatusOK, resp)
}
