package router

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/nathan/gopportunitiesbb/handler"
)

func initializeRoutes(router *gin.Engine)  {
	v1 := router.Group("/api/v1")
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
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.GET("/opening/:id", handler.ShowOpeningHandler)
		v1.DELETE("/opening/:id", handler.DeleteOpeningHandler)
		v1.PUT("/opening/:id", handler.UpdateOpeningHandler)
	}
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