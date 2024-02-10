package accounthandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *accountHandler) ChangeStatusAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Query("id")
		accountID, _ := strconv.ParseUint(idParam, 10, 64)

		statusParam := c.Query("status")
		status, _ := strconv.ParseUint(statusParam, 10, 64)
		err := hdl.accountUC.ChangeStatusAccount(c.Request.Context(), int(accountID), int(status))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
