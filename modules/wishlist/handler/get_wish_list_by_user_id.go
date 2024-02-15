package wishlisthandler

import (
	"net/http"
	"iTask/common"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *wishListHandler) GetWishListByUserID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		wishListID := ctx.Param("user_id")
		id, err := strconv.Atoi(wishListID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		var paging common.Paging

		page, _ := strconv.Atoi(ctx.Query("page"))
		limit, _ := strconv.Atoi(ctx.Query("limit"))

		paging.Page = page
		paging.Limit = limit

		res, err := hdl.wishListUC.GetWishListByUserID(ctx.Request.Context(), id, &paging)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": res, "paging": paging})

	}
}
