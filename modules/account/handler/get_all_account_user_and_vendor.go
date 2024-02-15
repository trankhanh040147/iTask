package accounthandler

import (
	"net/http"
	"iTask/common"
	"iTask/modules/account/convert"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *accountHandler) GetAllAccountUserAndVendor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging

		page, _ := strconv.Atoi(c.Query("page"))
		limit, _ := strconv.Atoi(c.Query("limit"))

		paging.Page = page
		paging.Limit = limit

		result, err := hdl.accountUC.GetAllAccountUserAndVendor(c.Request.Context(), &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		res := convert.ConvertAccountEntityToInfoMangageForAdmin(result)
		c.JSON(http.StatusOK, gin.H{"data": res, "paging": paging})
	}
}
