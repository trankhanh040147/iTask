package wishlisthandler

import (
	"net/http"
	wishlistiomodel "paradise-booking/modules/wishlist/iomodel"

	"github.com/gin-gonic/gin"
)

func (hdl *wishListHandler) CreateWishList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var wishListBody wishlistiomodel.CreateWishListReq

		if err := c.ShouldBind(&wishListBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		res, err := hdl.wishListUC.CreateWishList(c.Request.Context(), &wishListBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
