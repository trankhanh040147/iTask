package bookingratinghandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *bookingratinghandler) GetCommentByVendorID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		vendorID := ctx.Param("id")
		id, err := strconv.Atoi(vendorID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		res, err := hdl.placeRatingUC.GetCommentByVendorID(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": res})

	}
}
