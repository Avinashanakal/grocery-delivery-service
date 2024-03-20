package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Orders(ctx *gin.Context) {
	customerID := ctx.Param("customerID")
	id, _ := strconv.Atoi(customerID)

	if id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Customer ID is invalid",
			"id":    id,
		})
		return
	}

	orders, err := c.api.Orders(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}
