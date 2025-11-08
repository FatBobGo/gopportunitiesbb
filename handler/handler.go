package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func CreateAuthor(ctx *gin.Context) {
	name := ctx.Param("name")

	var requestBody struct {
		Content string `json:"content"`
		Created_at string `json:"created_at"`
		Modified_by string `json:"modified_by"`
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	
	response := gin.H {
		"name": name,
		"content": requestBody.Content,
		"created_at": requestBody.Created_at,
		"modified_by": requestBody.Modified_by,
	}

	ctx.JSON(http.StatusOK, response)

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"message": "Author " + name + " created successfully",
	// })
}

