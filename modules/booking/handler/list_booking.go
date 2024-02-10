package bookinghandler

import (
	"net/http"
	"paradise-booking/common"
	"paradise-booking/modules/booking/iomodel"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *bookingHandler) ListBooking() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requester := ctx.MustGet("Account").(common.Requester)

		var paging common.Paging
		var filter iomodel.FilterListBooking
		page, _ := strconv.Atoi(ctx.Query("page"))
		limit, _ := strconv.Atoi(ctx.Query("limit"))

		paging.Page = page
		paging.Limit = limit

		if err := ctx.ShouldBind(&filter); err != nil {
			panic(err)
		}

		res, err := hdl.bookingUC.ListBooking(ctx.Request.Context(), &paging, &filter, requester.GetID())
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, gin.H{"data": res, "paging": paging})
	}
}
