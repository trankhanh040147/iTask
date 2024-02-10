package accounthandler

import (
	"net/http"
	"paradise-booking/modules/account/iomodel"

	"github.com/gin-gonic/gin"
)

func (hdl *accountHandler) RegisterAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data iomodel.AccountRegister

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		result, err := hdl.accountUC.CreateAccount(c.Request.Context(), &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}
