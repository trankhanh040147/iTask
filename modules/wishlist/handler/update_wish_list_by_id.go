package wishlisthandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *wishListHandler) UpdateWishListByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		wishListID := ctx.Param("id")
		id, err := strconv.Atoi(wishListID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		title := ctx.Query("title")

		err = hdl.wishListUC.UpdateByID(ctx.Request.Context(), id, title)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": true})

	}
}
