package ginprojectmembers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	accountstorage "iTask/modules/account/storage"
	projectStore "iTask/modules/project/storage"
	"iTask/modules/project_members/biz"
	"iTask/modules/project_members/storage"
	verifyemailsstorage "iTask/modules/verify_emails/storage"
	"net/http"
)

func FindUninvitedMember(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var queryString struct {
			ProjectId int    `json:"project_id" form:"project_id"`
			Email     string `json:"email" form:"email"`
		}

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// dependency
		store := storage.NewSQLStore(db)
		accountSto := accountstorage.NewAccountStorage(db)
		projectSto := projectStore.NewSQLStore(db)
		emailSto := verifyemailsstorage.NewVerifyEmailsStorage(db)
		business := biz.NewFindUninvitedMemberBiz(store, accountSto, projectSto, emailSto)

		data, err := business.FindUninvitedMember(c, queryString.ProjectId, queryString.Email)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, data)
	}
}

// todo: add field is_member
