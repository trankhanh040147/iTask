package ginitem

import (
	"social-todo-list/common"
	"social-todo-list/module/item/biz"
	"social-todo-list/module/item/model"
	"social-todo-list/module/item/repository"
	"social-todo-list/module/item/storage"
	storage2 "social-todo-list/module/userlikeitem/storage"

	"net/http"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListItem(serviceCtx goservice.ServiceContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var queryString struct {
			common.Paging
			model.Filter
		}

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		queryString.Paging.Process()

		db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		store := storage.NewSQLStore(db)
		likeStore := storage2.NewSQLStore(db)
		repo := repository.NewListItemRepo(store, likeStore, requester)
		business := biz.NewListItemBiz(repo, requester)

		result, err := business.ListItem(c.Request.Context(), &queryString.Filter, &queryString.Paging)

		// SELECT * FROM todo_items ORDER BY id ASC LIMIT paging.Limit OFFSET (paging.Page - 1) * paging.Limit
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// fea_FakeID

		for i := range result {
			result[i].Mask()
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, queryString.Paging, queryString.Filter))
	}
}
