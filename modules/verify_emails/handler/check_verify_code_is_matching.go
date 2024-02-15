package verifyemailshanlder

import (
	"net/http"
	"iTask/constant"

	"github.com/gin-gonic/gin"
)

func (hdl *verifyEmailsHandler) CheckVerifyCodeIsMatching() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Query("email")
		verifyCode := c.Query("secret_code")

		isExpired, err := hdl.verifyEmailsUC.CheckVerifyCodeIsMatching(c.Request.Context(), email, verifyCode)
		if isExpired || err != nil {
			c.Redirect(http.StatusMovedPermanently, constant.UrlVerifyEmailFail)
			return
		}

		c.Redirect(http.StatusMovedPermanently, constant.UrlVerifyEmailSuccess)
		//c.JSON(http.StatusOK, gin.H{"status": true})
	}
}
