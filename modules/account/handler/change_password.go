package accounthandler

import (
	"net/http"
	"paradise-booking/common"
	"paradise-booking/modules/account/iomodel"

	"github.com/gin-gonic/gin"
)

func (hdl *accountHandler) ChangePassword() gin.HandlerFunc {
	return func(c *gin.Context) {

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		changePassModel := iomodel.ChangePassword{}
		if err := c.ShouldBind(&changePassModel); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		err := hdl.accountUC.ChangePassword(c.Request.Context(), requester.GetEmail(), &changePassModel)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
