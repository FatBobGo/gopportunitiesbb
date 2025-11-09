package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathan/gopportunitiesbb/configs"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	var err error
	db, err = config.GetSQLite()
	if err != nil {
		logger.Errorf("not able to init handler: %v", err)
	}
}

func GetStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// @BasePath /api/v1

// @Summary Create an opening job
// @Description Create a new job opening with role, company, and remote status
// @Tags openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
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

