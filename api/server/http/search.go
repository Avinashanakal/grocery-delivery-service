package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Search(ctx *gin.Context) {

	//used query parameter to search
	searchQuery := ctx.Query("q")

	searchResults, err := c.api.Search(ctx.Request.Context(), searchQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, searchResults)
}
