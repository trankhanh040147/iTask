package amenityhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (hdl *amenityHandler) GetAllConfigAmenity() gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := hdl.amenityUC.GetAllConfigAmenity(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})

	}
}
