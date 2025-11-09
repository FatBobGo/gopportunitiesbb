package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nathan/gopportunitiesbb/docs"
	_ "github.com/nathan/gopportunitiesbb/docs"
	"github.com/nathan/gopportunitiesbb/handler"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func initializeRoutes(router *gin.Engine)  {
	// Init handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.GET("/status", handler.GetStatus)
		
		v1.GET("/hello", func(c *gin.Context) {
			c.String(200, "Hello, World!")
		})

		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})

		v1.POST("/author/:name", handler.CreateAuthor)
		// Opening routes
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.GET("/opening/:id", handler.ShowOpeningHandler)
		v1.DELETE("/opening/:id", handler.DeleteOpeningHandler)
		v1.PUT("/opening/:id", handler.UpdateOpeningHandler)

		// v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func Initialize()  {
	router := gin.Default()
	initializeRoutes(router)

	log.Println("Server is running on port 8080")
	log.Println("Visit http://localhost:8080/hello to see the greeting")
	log.Println("Press Ctrl+C to stop the server")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}