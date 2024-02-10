package amenityhandler

import (
	"net/http"
	"paradise-booking/modules/amenity/iomodel"

	"github.com/gin-gonic/gin"
)

func (hdl *amenityHandler) DeleteAmenityByListID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var data iomodel.DeleteAmenityReq
		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		err := hdl.amenityUC.DeleteAmenityByListId(ctx, &data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}
