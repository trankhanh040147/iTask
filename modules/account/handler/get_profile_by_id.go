package accounthandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *accountHandler) GetAccountByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		accountID, _ := strconv.ParseUint(idParam, 10, 64)
		result, err := hdl.accountUC.GetAccountByID(c.Request.Context(), int(accountID))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}
