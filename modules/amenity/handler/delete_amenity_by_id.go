package amenityhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *amenityHandler) DeleteAmenityByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		amenityID := ctx.Param("id")
		id, err := strconv.Atoi(amenityID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		err = hdl.amenityUC.DeleteAmenityById(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}
