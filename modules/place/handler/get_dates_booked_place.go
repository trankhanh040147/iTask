package placehandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *placeHandler) GetDatesBookedPlace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		placeID := ctx.Query("place_id")
		id, err := strconv.Atoi(placeID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		dates, err := hdl.placeUC.GetDatesBookedPlace(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": dates})
	}
}
