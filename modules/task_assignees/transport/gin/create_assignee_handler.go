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

func CreateAssignee(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var AssigneeData model.TaskAssigneeCreation

		if err := c.ShouldBind(&AssigneeData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := storage.NewSQLStore(db)

		business := biz.NewCreateAssigneeBiz(store, requester)

		// step 3: use db.Create to
		if err := business.CreateNewAssignee(c.Request.Context(), &AssigneeData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// step 4: print data of the inserted record
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(AssigneeData.Id))
	}
}
