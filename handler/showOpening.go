package handler

import (
	"net/http"

	"github.com/MarcosViniicius/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary Show opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	// Find Opening by ID
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening with id not found ")
		return
	}
	sendSuccess(ctx, "select-opening", opening)
}
