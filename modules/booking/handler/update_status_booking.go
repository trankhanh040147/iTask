package bookinghandler

import (
	"net/http"
	"paradise-booking/constant"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *bookingHandler) UpdateStatusBooking() gin.HandlerFunc {
	return func(c *gin.Context) {
		bookingID, _ := c.GetQuery("booking_id")
		status, _ := c.GetQuery("status")

		bookingId, _ := strconv.Atoi(bookingID)
		statusInt, _ := strconv.Atoi(status)

		err := hdl.bookingUC.UpdateStatusBooking(c.Request.Context(), bookingId, statusInt)
		if err != nil {
			c.Redirect(http.StatusMovedPermanently, constant.UrlConfirmBookingFail)
			return
		}
		c.Redirect(http.StatusMovedPermanently, constant.UrlConfirmBookingSuccess)

		// c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
