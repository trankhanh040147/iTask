package ginproject

import (
	"iTask/common"
	"iTask/modules/project/biz"
	"iTask/modules/project/model"
	"iTask/modules/project/storage"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateProject(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data model.ProjectUpdate
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			// >> c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),}) is called. This sends a JSON response to the client with a status code of 400 (Bad Request). The JSON body of the response contains a single property error, which is set to the string representation of the error returned by ShouldBind. The gin.H function is a shortcut for creating a map in Go, which in this case is used to create the JSON object.
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// dependency
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		store := storage.NewSQLStore(db)
		business := biz.NewUpdateProjectBiz(store, requester)

		if err = business.UpdateProject(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
