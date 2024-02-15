package bookingratinghandler

import (
	"net/http"
	"paradise-booking/common"
	"paradise-booking/modules/booking_rating/iomodel"

	"github.com/gin-gonic/gin"
)

func (hdl *bookingratinghandler) MakeComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet("Account").(common.Requester)

		input := iomodel.CreateBookingRatingReq{}
		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		res, err := hdl.placeRatingUC.MakeComment(c.Request.Context(), requester.GetID(), &input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
