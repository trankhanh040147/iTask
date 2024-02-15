package accounthandler

import (
	"net/http"
	"paradise-booking/modules/account/iomodel"

	"github.com/gin-gonic/gin"
)

func (hdl *accountHandler) ResetPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updatePassModel iomodel.ResetPassword
		if err := c.ShouldBind(&updatePassModel); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		email := c.Query("email")
		err := hdl.accountUC.UpdatePassword(c.Request.Context(), email, updatePassModel.NewPassword)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
