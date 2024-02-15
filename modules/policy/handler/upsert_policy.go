package policieshandler

import (
	"net/http"
	"iTask/modules/policy/iomodel"

	"github.com/gin-gonic/gin"
)

func (hdl *policyHandler) UpsertPolicy() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input iomodel.CreatePolicyReq

		if err := ctx.ShouldBind(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		err := hdl.policyUC.UpSearchPolicy(ctx, &input)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}
