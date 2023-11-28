package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// UserController represents a user-related controller
type UserController struct{}

// GetUserInfo is a controller method to get user information
func (uc *UserController) GetUserInfo(c *gin.Context) {
	userID := c.Param("id")
	// Fetch user information from the database or other data source
	// For simplicity, we'll just return a JSON response.
	c.JSON(200, gin.H{"id": userID, "name": "John Doe", "email": "john@example.com"})
}

func main() {
	router := gin.Default()
	userController := &UserController{} //creando un controlador

	// Ruta para los archivos estáticos de producción de Vite
	router.Static("/data", "http://localhost:5173")

	// // Ruta principal que sirve el archivo index.html
	// router.GET("/data", func(c *gin.Context) {
	// 	c.File("./path/to/your/vue/app/dist/index.html")
	// })

	router.Use(cors.Default())

	// Basic route
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Bill App")
	})

	// Route with URL parameters
	// router.GET("/users/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	c.String(200, "User ID: "+id)
	// })

	// Route with query parameters
	// router.GET("/search", func(c *gin.Context) {
	// 	query := c.DefaultQuery("q", "default-value")
	// 	c.String(200, "Search query: "+query)
	// })

	// Route using the UserController
	router.GET("/data", userController.GetUserInfo)

	router.Run(":8080")
}
