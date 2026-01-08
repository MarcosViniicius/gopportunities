package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func ShowOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET opening",
	})
}

func CreateOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "post opening",
	})
}
func DeleteOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "delete opening",
	})
}
func UpdateOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "put opening",
	})
}
func ListOpeningsHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "list opening",
	})
}