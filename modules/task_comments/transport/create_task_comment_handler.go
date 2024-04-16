package gintaskcomments

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"iTask/common"
	"iTask/modules/task_comments/biz"
	"iTask/modules/task_comments/model"
	"iTask/modules/task_comments/storage"

	"net/http"
)

func CreateTaskComment(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var TaskCommentData model.TaskCommentCreation

		if err := c.ShouldBind(&TaskCommentData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		store := storage.NewSQLStore(db)
		business := biz.NewCreateTaskCommentBiz(store, requester)

		TaskCommentData.CreatedBy = requester.GetUserId()

		// step 3: use db.Create to
		if err := business.CreateNewTaskComment(c.Request.Context(), &TaskCommentData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// step 4: print data of the inserted record
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(TaskCommentData.Id))
	}
}
