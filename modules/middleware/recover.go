package middleware

import (
	"iTask/common"

	"github.com/gin-gonic/gin"
)

func (m *middlewareManager) Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.Header("Content-Type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					ctx.AbortWithStatusJSON(appErr.StatusCode, appErr)
					// call function recover of gin
					panic(err)
				}

				appErr := common.ErrInternal(err.(error))
				ctx.AbortWithStatusJSON(appErr.StatusCode, appErr)
				// call function recover of gin
				panic(err)
			}
		}()
		ctx.Next()
	}
}
