package gintask

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"iTask/common"
	"iTask/modules/task/biz"
	"iTask/modules/task/model"
	"iTask/modules/task/storage"
	"net/http"
)

func CreateTask(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var TaskData model.TaskCreation

		if err := c.ShouldBind(&TaskData); err != nil {
			// >> c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),}) is called. This sends a JSON response to the client with a status code of 400 (Bad Request). The JSON body of the response contains a single property error, which is set to the string representation of the error returned by ShouldBind. The gin.H function is a shortcut for creating a map in Go, which in this case is used to create the JSON object.
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		TaskData.CreatedBy = requester.GetUserId()

		store := storage.NewSQLStore(db)

		business := biz.NewCreateTaskBiz(store, requester)

		// step 3: use db.Create to
		if err := business.CreateNewTask(c.Request.Context(), &TaskData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// step 4: print data of the inserted record
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(TaskData.Id))
	}
}
