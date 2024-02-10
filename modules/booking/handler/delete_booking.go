package bookinghandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *bookingHandler) DeleteBookingByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bookingID := ctx.Param("id")
		id, err := strconv.Atoi(bookingID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		err = hdl.bookingUC.DeleteBookingByID(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}
