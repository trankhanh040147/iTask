package placehandler

import (
	"net/http"
	"paradise-booking/common"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *placeHandler) DeletePlaceByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		requester := ctx.MustGet("Account").(common.Requester)

		placeID := ctx.Query("id")
		id, err := strconv.Atoi(placeID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		err = hdl.placeUC.DeletePlaceByID(ctx.Request.Context(), id, requester.GetEmail())
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}
