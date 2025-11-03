package router

import (
	"log"
	"github.com/gin-gonic/gin"
)

func Initialize()  {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	router.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Println("Server is running on port 8080")
	log.Println("Visit http://localhost:8080/hello to see the greeting")
	log.Println("Press Ctrl+C to stop the server")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}