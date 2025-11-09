package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nathan/gopportunitiesbb/schemas"
)

func ShowOpeningHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		sendErrorResponse(ctx, http.StatusBadRequest, "the URL param ID should be integer")
		return
	}

	var opening schemas.Opening
	result := db.First(&opening, id)
	if result.Error != nil {
		sendErrorResponse(ctx, http.StatusNotFound, "Failed to find opening with id " + idStr)
		return
	}
	sendSuccess(ctx, "show-opening", opening)
}