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

// Create an opening job
// @Summary Create an opening job
// @Produce json
// @Param role body string false "job role" minlength(3) maxlength(100)
// @Param company body string false "job description" Enums(0, 1) default(1)
// @Param Remote body bool true "Remote job ind" boolean
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{} "参数错误"
// @Failure 500 {object} string "内部错误"
// @Router /api/v1/opening [post]
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

