package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Printf("hohohooho Request - Method: %s | Status: %d | Duration: %v", c.Request.Method, c.Writer.Status(), duration)
	}
}

func AuthMiddleware() gin.HandlerFunc {
	// In a real-world application, you would perform proper authentication here.
	// For the sake of this example, we'll just check if an API key is present.
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		if apiKey == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}

func main() {
	// Your Gin application code goes here
	// Create a new Gin router
	router := gin.Default()

	// Use our custom authentication middleware for a specific group of routes
	authGroup := router.Group("/api") // se pueden agrupar rutas
	authGroup.Use(AuthMiddleware())
	{
		authGroup.GET("/data", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Authenticated and authorized!"})
		})
	}
	router.Use(LoggerMiddleware())

	// Define a route for the root URL
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	router.GET("/bye", func(c *gin.Context) {
		c.String(200, "Bye, World!")
	})

	router.GET("/prueba", func(c *gin.Context) {
		c.String(200, "tararararara")
	})

	// Run the server on port 8080
	router.Run("mi-host:8080")
}
