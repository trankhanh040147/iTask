package placewishlisthandler

import (
	"net/http"
	"paradise-booking/common"
	"paradise-booking/modules/place_wishlist/iomodel"

	"github.com/gin-gonic/gin"
)

func (hdl *placeWishListHandler) CreatePlaceWishList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var placeWishListBody iomodel.CreatePlaceWishListReq

		if err := c.ShouldBind(&placeWishListBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		requester := c.MustGet("Account").(common.Requester)
		res, err := hdl.placeWishListUC.CreatePlaceWishList(c.Request.Context(), &placeWishListBody, requester.GetID())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
