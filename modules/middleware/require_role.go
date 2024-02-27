package middleware

import (
	"errors"
	"iTask/common"
	"iTask/entities"

	"github.com/gin-gonic/gin"
)

func (m *middlewareManager) RequiredRoles(roles ...entities.UserRole) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		account := ctx.MustGet(common.CurrentUser).(*entities.Account)
		for i := range roles {
			if int(account.Role) == int(roles[i]) {
				ctx.Next()
				return
			}
		}
		//ctx.JSON(http.StatusForbidden, gin.H{"error": "you have no permission"})
		panic(common.ErrForbidden(errors.New("have no permission")))
	}
}
