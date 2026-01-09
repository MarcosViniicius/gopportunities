package router

import (
	docs "github.com/MarcosViniicius/gopportunities/docs"
	"github.com/MarcosViniicius/gopportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {

	// Initialize handler
	handler.InitializeHandler()
	BasePath:= "/api/v1"
	docs.SwaggerInfo.BasePath = BasePath
	v1:= router.Group(BasePath)
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler,)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings",handler.ListOpeningsHandler)
	}

	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}