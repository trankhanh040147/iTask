package ginprojectmembers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"iTask/modules/project_members/biz"
	"iTask/modules/project_members/storage"
	taskStorage "iTask/modules/task/storage"
	taskAssigneeStorage "iTask/modules/task_assignees/storage"
	"net/http"
)

func ListUnassignedMembers(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var queryString struct {
			ProjectId int `json:"project_id" form:"project_id"`
			TaskId    int `json:"task_id" form:"task_id"`
		}

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// dependency
		store := storage.NewSQLStore(db)
		taskAssigneeStore := taskAssigneeStorage.NewSQLStore(db)
		taskStore := taskStorage.NewSQLStore(db)
		business := biz.NewListUnassignedMembersBiz(store, taskAssigneeStore, taskStore)

		data, err := business.ListUnassignedMembers(c.Request.Context(), queryString.ProjectId, queryString.TaskId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, data)
	}
}

// todo: add field is_member
