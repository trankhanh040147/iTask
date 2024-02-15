package bookinghandler

import (
	"net/http"
	"iTask/common"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *bookingHandler) GetBookingByPlaceID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		placeID := ctx.Query("place_id")
		id, err := strconv.Atoi(placeID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		var paging common.Paging

		page, _ := strconv.Atoi(ctx.Query("page"))
		limit, _ := strconv.Atoi(ctx.Query("limit"))

		paging.Page = page
		paging.Limit = limit

		res, err := hdl.bookingUC.GetBookingByPlaceID(ctx, id, &paging)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": res, "paging": paging})
	}
}
