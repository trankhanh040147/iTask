package gintask

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"iTask/common"
	"iTask/modules/task/biz"
	"iTask/modules/task/model"
	repository "iTask/modules/task/repo"
	"iTask/modules/task/storage"
	"net/http"
)

func ListTask(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var queryString struct {
			common.Paging
			model.Filter
		}

		// Default value for CreatedDayRange
		// queryString.CreatedDayRange = -1
		//
		//if err := c.ShouldBind(&queryString); err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{
		//		"error": err.Error(),
		//	})
		//	return
		//}
		//
		//queryString.Paging.Process()

		//requester := c.MustGet(common.CurrentUser).(common.Requester)
		//requester := nil
		store := storage.NewSQLStore(db)
		//taskStore := taskStorage.NewSQLStore(db)
		repo := repository.NewListTaskRepo(store, nil)
		business := biz.NewListTaskBiz(repo, nil)

		result, err := business.ListTask(c.Request.Context(), &queryString.Filter, &queryString.Paging)

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
