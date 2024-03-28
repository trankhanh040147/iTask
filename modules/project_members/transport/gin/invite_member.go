package ginprojectmembers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	accountstorage "iTask/modules/account/storage"
	projectStore "iTask/modules/project/storage"
	"iTask/modules/project_members/biz"
	"iTask/modules/project_members/model"
	"iTask/modules/project_members/storage"
	"iTask/worker"
	"net/http"
)

func InviteMember(db *gorm.DB, taskDistributor worker.TaskDistributor) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data model.ProjectMemberInvitation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// dependency
		store := storage.NewSQLStore(db)
		accountSto := accountstorage.NewAccountStorage(db)
		projectSto := projectStore.NewSQLStore(db)
		business := biz.NewInviteMemberBiz(store, accountSto, projectSto, taskDistributor)

		if err := business.InviteMember(c, data.UserEmail, data.ProjectId); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, true)
	}
}
