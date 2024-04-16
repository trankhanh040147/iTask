package ginprojectmembers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"iTask/common"
	"iTask/modules/project_members/biz"
	"iTask/modules/project_members/storage"
	"net/http"
)

func DeleteMember(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		var queryString struct {
			UserId    int `json:"user_id" form:"user_id"`
			ProjectId int `json:"project_id" form:"project_id"`
		}

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewDeleteMemberBiz(store)

		if err := business.DeleteMember(c.Request.Context(), queryString.UserId, queryString.ProjectId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
