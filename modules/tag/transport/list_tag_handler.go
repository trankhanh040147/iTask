package gintag

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"iTask/common"
	"iTask/modules/tag/biz"
	"iTask/modules/tag/model"
	"iTask/modules/tag/storage"
	"net/http"
)

func ListTag(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var queryString struct {
			common.Paging
			model.Filter
		}

		// Default value for CreatedDayRange
		//queryString.CreatedDayRange = -1

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		queryString.Paging.Process()

		//requester := c.MustGet(common.CurrentUser).(common.Requester)
		//requester := nil
		store := storage.NewSQLStore(db)
		business := biz.NewListTagBiz(store)

		result, err := business.ListTag(c.Request.Context(), &queryString.Filter, &queryString.Paging)

		// SELECT * FROM todo_items ORDER BY id ASC LIMIT paging.Limit OFFSET (paging.Page - 1) * paging.Limit
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// fea_FakeID

		// for i := range result {
		// 	result[i].Mask()
		// }

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, queryString.Paging, queryString.Filter))
	}
}
