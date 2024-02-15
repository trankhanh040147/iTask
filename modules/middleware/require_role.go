package middleware

import (
	"errors"
	"paradise-booking/common"
	"paradise-booking/constant"
	"paradise-booking/entities"

	"github.com/gin-gonic/gin"
)

func (m *middlewareManager) RequiredRoles(roles ...constant.Role) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		account := ctx.MustGet("Account").(*entities.Account)
		for i := range roles {
			if account.Role == int(roles[i]) {
				ctx.Next()
				return
			}
		}
		//ctx.JSON(http.StatusForbidden, gin.H{"error": "you have no permission"})
		panic(common.ErrForbidden(errors.New("have no permission")))
	}
}
