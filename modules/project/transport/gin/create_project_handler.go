package ginproject

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"iTask/common"
	"iTask/modules/project/biz"
	"iTask/modules/project/model"
	"iTask/modules/project/storage"
	storage3 "iTask/modules/project_tags/storage"
	storage2 "iTask/modules/tag/storage"
	"net/http"
)

func CreateProject(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var projectData model.ProjectCreation

		if err := c.ShouldBind(&projectData); err != nil {
			// >> c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),}) is called. This sends a JSON response to the client with a status code of 400 (Bad Request). The JSON body of the response contains a single property error, which is set to the string representation of the error returned by ShouldBind. The gin.H function is a shortcut for creating a map in Go, which in this case is used to create the JSON object.
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		projectData.CreatedBy = requester.GetUserId()

		store := storage.NewSQLStore(db)
		tagStore := storage2.NewSQLStore(db)
		projectTagStore := storage3.NewSQLStore(db)

		business := biz.NewCreateProjectBiz(store, projectTagStore, tagStore, requester)

		// step 3: use db.Create to
		if err := business.CreateNewProject(c.Request.Context(), &projectData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// step 4: print data of the inserted record
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(projectData.Id))
	}
}
