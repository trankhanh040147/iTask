package accounthandler

import (
	"net/http"
	"paradise-booking/modules/account/iomodel"

	"github.com/gin-gonic/gin"
)

func (hdl *accountHandler) LoginAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var accountLogin iomodel.AccountLogin

		if err := c.ShouldBind(&accountLogin); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		result, err := hdl.accountUC.LoginAccount(c.Request.Context(), &accountLogin)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}
