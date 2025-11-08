package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nathan/gopportunitiesbb/config"
)

func CreateOpeningHandler(ctx *gin.Context) {
	logger := config.GetLogger("CreateOpeningHandler")

	request := struct{
		Role string `json:"role"`
	}{}

	ctx.BindJSON(&request)
	logger.Infof("Received request to create opening: %+v", request)

	// request := CreateOpeningRequest{}

	// ctx.BindJSON(&request)

	// if err := db.Create(&request).Error; err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "Failed to create opening",
	// 	})
	// 	return
	// }

	// response := OpeningResponse{
	// 	ID:        request.ID,
	// 	CreatedAt: request.CreatedAt,
	// 	UpdatedAt: request.UpdatedAt,
	// 	Role:      request.Role,
	// 	Company:   request.Company,
	// 	Location:  request.Location,
	// 	Remote:    request.Remote,
	// 	Link:      request.Link,
	// 	Salary:    request.Salary,
	// }

	// ctx.JSON(http.StatusCreated, response)

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"message": "POST Opening",
	// })
}