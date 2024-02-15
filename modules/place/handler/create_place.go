package placehandler

import (
	"net/http"
	"paradise-booking/common"
	placeIomodel "paradise-booking/modules/place/iomodel"

	"github.com/gin-gonic/gin"
)

func (hdl *placeHandler) CreatePlace() gin.HandlerFunc {
	return func(c *gin.Context) {
		var placeBody placeIomodel.CreatePlaceReq

		requester := c.MustGet("Account").(common.Requester)
		if err := c.ShouldBind(&placeBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		err := hdl.placeUC.CreatePlace(c.Request.Context(), &placeBody, requester.GetEmail())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
