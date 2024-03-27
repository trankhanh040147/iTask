package gintaskassignee

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"iTask/common"
	"iTask/modules/task_assignees/biz"
	"iTask/modules/task_assignees/model"
	"iTask/modules/task_assignees/storage"
	"net/http"
)

func ListAssignee(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var queryString struct {
			model.Filter
		}

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewListAssigneeBiz(store)

		result, err := business.ListAssignee(c.Request.Context(), &queryString.Filter)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, nil, queryString.Filter))
	}
}
