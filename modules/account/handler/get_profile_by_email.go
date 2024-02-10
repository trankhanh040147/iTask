package accounthandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (hdl *accountHandler) GetAccountByEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Query("email")
		result, err := hdl.accountUC.GetAccountByEmail(c.Request.Context(), email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}
