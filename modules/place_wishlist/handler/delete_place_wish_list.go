package placewishlisthandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *placeWishListHandler) DeletePlaceWishList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_placeID := ctx.Param("place_id")
		placeID, err := strconv.Atoi(_placeID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		_wishListID := ctx.Param("wishlist_id")
		wishListID, err := strconv.Atoi(_wishListID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		err = hdl.placeWishListUC.DeletePlaceWishList(ctx.Request.Context(), placeID, wishListID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": true})

	}
}
