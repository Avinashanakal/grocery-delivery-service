package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Avinashanakal/grocery-delivery-service/internal/pkg/groceries"
)

func (c *Controller) Groceries(ctx *gin.Context) {
	var groceryRequest groceries.GroceryRequest

	if err := ctx.ShouldBind(&groceryRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	err := c.api.Groceries(ctx.Request.Context(), groceryRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Groceries Added to Inventory",
	})

}
