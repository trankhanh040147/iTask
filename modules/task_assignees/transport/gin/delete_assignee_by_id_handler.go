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

func DeleteAssignee(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		var queryString struct {
			model.TaskAssigneeDeletion
		}

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewDeleteAssigneeBiz(store)

		if err := business.DeleteAssignee(c.Request.Context(), queryString.UserId, queryString.TaskId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
