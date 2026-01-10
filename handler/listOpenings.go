package handler

import (
	"net/http"

	"github.com/MarcosViniicius/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary      List openings
// @Description  List all job openings
// @Tags         openings
// @Accept       json
// @Produce      json
// @Success      200  {array}   schemas.Opening
// @Failure      500  {object}  ErrorResponse
// @Router       /openings [get]

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	// Delete opening
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(ctx, "list-openings", openings)
}
