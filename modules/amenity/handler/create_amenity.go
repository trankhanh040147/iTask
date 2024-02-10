package amenityhandler

import (
	"net/http"
	"paradise-booking/modules/amenity/iomodel"

	"github.com/gin-gonic/gin"
)

func (hdl *amenityHandler) CreateAmenity() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data iomodel.CreateAmenityReq
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		err := hdl.amenityUC.CreateAmenity(c.Request.Context(), &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
