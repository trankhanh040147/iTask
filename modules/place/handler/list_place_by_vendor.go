package placehandler

import (
	"net/http"
	"paradise-booking/common"

	"github.com/gin-gonic/gin"
)

func (hdl *placeHandler) ListPlaceByVendor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		places, err := hdl.placeUC.ListPlaceByVendor(ctx.Request.Context(), requester.GetEmail())
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, gin.H{"data": places})
	}
}
