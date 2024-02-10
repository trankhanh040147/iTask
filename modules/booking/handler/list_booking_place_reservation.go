package bookinghandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *bookingHandler) ListBookingNotReservation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		vendorID, err := strconv.Atoi(ctx.Query("vendor_id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		placeID, err := strconv.Atoi(ctx.Query("place_id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		res, err := hdl.bookingUC.ListPlaceReservationByVendor(ctx.Request.Context(), vendorID, placeID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": res})
	}
}
