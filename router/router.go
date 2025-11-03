package router

import (
	"log"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine)  {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/status", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"status": "ok",
			})
		})

		v1.GET("/hello", func(c *gin.Context) {
			c.String(200, "Hello, World!")
		})

		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})
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