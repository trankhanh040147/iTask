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

		// Default value for Filter fields
		// queryString.CreatedDayRange = -1
		queryString.Status = -1

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		queryString.Paging.Process()

		store := storage.NewSQLStore(db)
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

		//for _, t := range &result {
		//	(t).Parsing()
		//}
		//fmt.Printf("Result: %+v\n", result)

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, queryString.Paging, queryString.Filter))
	}
}
