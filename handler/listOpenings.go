package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathan/gopportunitiesbb/schemas"
)

func ListOpeningsHandler(ctx *gin.Context) {
	var openings []schemas.Opening
	result := db.Find(&openings)
	if result.Error != nil {
		sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to list openings")
		return
	}
	sendSuccess(ctx, "list-openings", openings)

}