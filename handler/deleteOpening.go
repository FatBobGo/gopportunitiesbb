package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nathan/gopportunitiesbb/schemas"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	// get the ID from the URL parameter
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		sendErrorResponse(ctx, http.StatusBadRequest, "the URL param ID should be integer")
		return
	}
	var opening schemas.Opening
	result := db.First(&opening, id)
	if result.Error != nil {
		sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to find opening with id " + idStr)
		return
	}
	result = db.Delete(&opening)
	if result.Error != nil {
		sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete opening with id " + idStr)
		return
	}
	if result.RowsAffected == 0 {
		sendErrorResponse(ctx, http.StatusNotFound, "Opening not deleted")
		return
	}
	sendSuccess(ctx, "delete-opening", opening)

}