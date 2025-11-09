package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathan/gopportunitiesbb/config"
	"github.com/nathan/gopportunitiesbb/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {
	logger := config.GetLogger("CreateOpeningHandler")

	// request := struct{
	// 	Role string `json:"role"`
	// }{}

	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)
	logger.Infof("Received request to create opening: %+v", request)

	if err := request.Validate(); err != nil {
		// logger.Errorf("Invalid request to create opening: %v", err.Error())
		// ctx.JSON(http.StatusBadRequest, gin.H{
		// 	"error": err.Error(),
		// })
		sendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	// request := CreateOpeningRequest{}

	// ctx.BindJSON(&request)

	opening := schemas.Opening {
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Failed to create opening in DB: %v", err)
		sendErrorResponse(ctx, http.StatusBadRequest, "Failed to create opening in database")
		return
	}
	// sendSuccessResponse(ctx, http.StatusAccepted, request)
	sendSuccess(ctx, "create-opening", opening)

}