package placehandler

import (
	"net/http"
	"paradise-booking/common"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *placeHandler) ListPlaceByVendorID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := ctx.Param("vendor_id")
		vendorID, err := strconv.Atoi(param)
		if err != nil {
			panic(err)
		}

		var paging common.Paging

		page, _ := strconv.Atoi(ctx.Query("page"))
		limit, _ := strconv.Atoi(ctx.Query("limit"))

		paging.Page = page
		paging.Limit = limit

		places, err := hdl.placeUC.ListPlaceByVendorByID(ctx.Request.Context(), vendorID, &paging)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, gin.H{"data": places, "paging": paging})
	}
}
