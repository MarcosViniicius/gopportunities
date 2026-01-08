package router

import "github.com/gin-gonic/gin"

// Exports the InitializeRouter function using the default Gin configuration.
func Initialize() {

	router := gin.Default()

	// Initialize Router
	initializeRoutes(router)

	
	router.Run() // Listen and serve on 0.0.0.0:8080
}