package verifyemailshanlder

import (
	"github.com/gin-gonic/gin"
	"iTask/constant"
	"net/http"
	"strconv"
)

func (hdl *verifyEmailsHandler) CheckProjectInvitation() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Query("email")
		verifyCode := c.Query("secret_code")
		projectId, err := strconv.Atoi(c.Query("project_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		isExpire, err := hdl.verifyEmailsUC.CheckProjectInvitation(c, email, verifyCode, projectId)
		if isExpire || err != nil {
			c.Redirect(http.StatusMovedPermanently, constant.UrlVerifyEmailFail)
			return
		}

		c.Redirect(http.StatusMovedPermanently, constant.UrlVerifyEmailSuccess)

	}
}
