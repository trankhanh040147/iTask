package ginprojectmembers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"iTask/common"
	"iTask/modules/project_members/biz"
	"iTask/modules/project_members/storage"
	"net/http"
	"strconv"
)

func ListMembersById(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("project_id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// dependency
		store := storage.NewSQLStore(db)
		business := biz.NewListMembersByIdBiz(store)

		data, err := business.ListMembersById(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
