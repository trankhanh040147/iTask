package wishlisthandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *wishListHandler) GetWishListByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		wishListID := ctx.Param("id")
		id, err := strconv.Atoi(wishListID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		res, err := hdl.wishListUC.GetWishListByID(ctx.Request.Context(), id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": res})

	}
}
