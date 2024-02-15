package bookinghandler

import (
	"net/http"
	"paradise-booking/modules/booking/iomodel"

	"github.com/gin-gonic/gin"
)

func (hdl *bookingHandler) CreateBooking() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data iomodel.CreateBookingReq
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		res, err := hdl.bookingUC.CreateBooking(c.Request.Context(), &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
