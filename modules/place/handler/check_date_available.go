package placehandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *placeHandler) CheckDateBookingAvailable() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dateFrom := ctx.Query("date_from")
		dateTo := ctx.Query("date_to")

		placeID := ctx.Query("place_id")
		id, err := strconv.Atoi(placeID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		isValid, err := hdl.placeUC.CheckDateBookingAvailable(ctx.Request.Context(), int64(id), dateFrom, dateTo)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": isValid})
	}
}
