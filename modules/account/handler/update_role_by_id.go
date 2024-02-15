package accounthandler

import (
	"net/http"
	"iTask/modules/account/iomodel"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *accountHandler) UpdateAccountRoleByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data iomodel.AccountChangeRole

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		id := c.Param("id")
		accountID, _ := strconv.ParseUint(id, 10, 64)
		err := hdl.accountUC.UpdateAccountRoleByID(c, &data, int(accountID))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
