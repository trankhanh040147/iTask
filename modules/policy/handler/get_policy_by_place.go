package policieshandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *policyHandler) GetPolicyByPlaceId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		placeId := ctx.Param("place_id")
		id, _ := strconv.Atoi(placeId)

		res, err := hdl.policyUC.GetPolicyByPlaceID(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": res})
	}
}
