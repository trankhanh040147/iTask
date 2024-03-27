package ginprojectmembers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	accountstorage "iTask/modules/account/storage"
	"iTask/modules/project_members/biz"
	"iTask/modules/project_members/model"
	"iTask/modules/project_members/storage"
	"net/http"
)

func InviteMember(db *gorm.DB) func(ctx *gin.Context) {
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
		business := biz.NewInviteMemberBiz(store, accountSto)

		if err := business.InviteMember(c, data.UserEmail, data.ProjectId); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, true)
	}
}
