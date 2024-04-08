package gintaskcomments

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"iTask/common"
	"iTask/modules/task_comments/biz"
	"iTask/modules/task_comments/storage"
	"net/http"
	"strconv"
)

func ListTaskCommentsByTaskId(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		taskId, err := strconv.Atoi(c.Param("task_id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var queryString struct {
			common.Paging
		}

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		queryString.Paging.Process()

		store := storage.NewSQLStore(db)
		business := biz.NewListTaskCommentBiz(store)

		result, err := business.ListTaskCommentsByTaskId(c.Request.Context(), taskId, &queryString.Paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, queryString.Paging, nil))

	}
}
