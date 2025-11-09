package handler

import "github.com/gin-gonic/gin"

func sendErrorResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, gin.H{
		"message": message,
		"errorCode": statusCode,
	})
}

// func sendErrorResponse(ctx *gin.Context, statusCode int, message string) {
// 	ctx.JSON(statusCode, gin.H{
// 		"error": message,
// 	})
// }