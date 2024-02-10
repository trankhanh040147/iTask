package placehandler

import (
	"net/http"
	"paradise-booking/common"
	placeIomodel "paradise-booking/modules/place/iomodel"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *placeHandler) UpdatePlace() gin.HandlerFunc {
	return func(c *gin.Context) {
		var placeBody placeIomodel.UpdatePlaceReq

		requester := c.MustGet("Account").(common.Requester)
		if err := c.ShouldBind(&placeBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		placeID, err := strconv.Atoi(c.Query("place_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		err = hdl.placeUC.UpdatePlace(c.Request.Context(), &placeBody, placeID, requester.GetEmail())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
