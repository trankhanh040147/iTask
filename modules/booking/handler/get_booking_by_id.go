package bookinghandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *bookingHandler) GetBookingByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bookingID := ctx.Param("id")
		id, err := strconv.Atoi(bookingID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		res, err := hdl.bookingUC.GetBookingByID(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}
