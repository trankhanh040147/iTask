package uploadhandler

import (
	"net/http"
	"iTask/common"

	"github.com/gin-gonic/gin"
)

func (hdl *uploadHandler) UploadFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		img, err := hdl.uploadUC.UploadFile(c.Request.Context(), fileHeader)
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		c.JSON(http.StatusOK, gin.H{"data": img})
	}
}
