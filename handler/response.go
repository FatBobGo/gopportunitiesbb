package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathan/gopportunitiesbb/schemas"
)

func sendErrorResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.Header("content-type", "application/json")
	ctx.JSON(statusCode, gin.H{
		"message": message,
		"errorCode": statusCode,
	})
}

// func sendSuccessResponse(ctx *gin.Context, statusCode int, opening CreateOpeningRequest) {
// 	ctx.Header("content-type", "application/json")
// 	ctx.JSON(statusCode, gin.H {
// 		"data" : opening,
// 	})
// }

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successful", op),
		"data": data,
	})
}

type ErrorResponse struct {
	Message string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOpeningResponse struct {
	Message string `json:"message"`
	Data schemas.OpeningResponse `json:"data"`
}