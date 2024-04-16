package ginprojectmembers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"iTask/common"
	storage2 "iTask/modules/project/storage"
	"iTask/modules/project_members/biz"
	"iTask/modules/project_members/storage"
	"net/http"
)

func UpdateMemberRole(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		var data struct {
			Role      *int `json:"role" column:"role"`
			ProjectId int  `json:"project_id"`
			MemberId  int  `json:"member_id"`
		}

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := storage.NewSQLStore(db)
		projectStore := storage2.NewSQLStore(db)
		business := biz.NewUpdateMemberRoleBiz(store, projectStore)

		if err := business.UpdateMemberRole(c.Request.Context(), data.ProjectId, data.MemberId, *data.Role); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
