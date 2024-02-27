package ginproject

import (
	"iTask/common"
	"iTask/modules/project/biz"
	"iTask/modules/project/storage"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// dependency
		store := storage.NewSQLStore(db)
		business := biz.NewGetProjectBiz(store)

		data, err := business.GetProjectById(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
