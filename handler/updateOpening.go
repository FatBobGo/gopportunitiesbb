package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nathan/gopportunitiesbb/schemas"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		sendErrorResponse(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		sendErrorResponse(ctx, http.StatusBadRequest, "the URL param ID should be integer")
		return
	}

	var updateRequest UpdateOpeningRequest
	ctx.ShouldBindJSON(&updateRequest)

	if err := updateRequest.Validate(); err != nil {
		sendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var opening schemas.Opening
	result := db.First(&opening, id)
	if result.Error != nil {
		sendErrorResponse(ctx, http.StatusNotFound, "Not able to update. Failed to find opening with id " + idStr)
		return
	}

	// 构建更新字段的 map
	updates := make(map[string]interface{})
	if updateRequest.Role != nil {
		updates["role"] = updateRequest.Role
	}
	if updateRequest.Company != nil {
		updates["company"] = updateRequest.Company
	}
	if updateRequest.Location != nil {
		updates["location"] = updateRequest.Location
	}
	if updateRequest.Remote != nil {
		updates["remote"] = updateRequest.Remote
	}
	if updateRequest.Link != nil {
		updates["link"] = updateRequest.Link
	}
	if updateRequest.Salary != nil {
		updates["salary"] = updateRequest.Salary
	}

	// 如果没有提供任何字段，返回错误
	if len(updates) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No fields provided for update"})
		return
	}
	// 执行更新
	if err := db.Model(&opening).Updates(updates).Error; err != nil {
		logger.Errorf("Failed to update opening in DB: %v", err)
		sendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	sendSuccess(ctx, "update-opening", opening)
}