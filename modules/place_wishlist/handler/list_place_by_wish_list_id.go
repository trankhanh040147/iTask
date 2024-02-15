package placewishlisthandler

import (
	"net/http"
	"paradise-booking/common"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *placeWishListHandler) ListPlaceByWishListID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		_wishListID := ctx.Query("wish_list_id")
		wishListID, err := strconv.Atoi(_wishListID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		var paging common.Paging

		page, _ := strconv.Atoi(ctx.Query("page"))
		limit, _ := strconv.Atoi(ctx.Query("limit"))

		paging.Page = page
		paging.Limit = limit

		res, err := hdl.placeWishListUC.GetPlaceByWishListID(ctx.Request.Context(), wishListID, &paging, requester.GetID())
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": res, "paging": paging})

	}
}
