package accounthandler

import (
	"net/http"
	"iTask/modules/account/iomodel"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *accountHandler) UpdatePersonalInfoAccountById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data iomodel.AccountUpdatePersonalInfo

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		id := c.Param("id")
		newId, _ := strconv.ParseUint(id, 10, 64)
		err := hdl.accountUC.UpdatePersonalInforAccountById(c, &data, int(newId))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
