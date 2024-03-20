package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Avinashanakal/grocery-delivery-service/internal/pkg/groceries"
)

func (c *Controller) Buy(ctx *gin.Context) {

	var orderRequest groceries.OrderRequest
	if err := ctx.ShouldBind(&orderRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	order, err := c.api.Buy(ctx.Request.Context(), orderRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, order)

}
