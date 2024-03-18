package ginproject

import (
	"iTask/common"
	"iTask/modules/project/biz"
	"iTask/modules/project/storage"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProject(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// dependency
		store := storage.NewSQLStore(db)
		business := biz.NewGetProjectBiz(store)

		data, err := business.GetProjectById(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		data.PriorityValue = data.Priority.String()
		data.StatusValue = data.Status.String()
		data.PrivacyValue = data.Privacy.String()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
