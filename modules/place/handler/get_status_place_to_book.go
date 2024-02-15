package placehandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *placeHandler) GetStatusPlaceToBook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dateFrom := ctx.Query("date_from")
		dateTo := ctx.Query("date_to")

		placeID := ctx.Query("place_id")
		id, err := strconv.Atoi(placeID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		res, err := hdl.placeUC.GetStatusPlaceToBook(ctx, id, dateFrom, dateTo)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": res})

	}
}
