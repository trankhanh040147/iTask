package amenityhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *amenityHandler) ListAmenityByPlaceID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		placeId := ctx.Param("place_id")
		id, err := strconv.Atoi(placeId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		res, err := hdl.amenityUC.ListAmenityByPlaceID(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": res})
	}
}
