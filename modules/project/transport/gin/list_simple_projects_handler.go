package ginproject

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"iTask/common"
	"iTask/modules/project/biz"
	"iTask/modules/project/model"
	repository "iTask/modules/project/repo"
	"iTask/modules/project/storage"
	storage2 "iTask/modules/project_members/storage"
	projecTagStorage "iTask/modules/project_tags/storage"
	taskStorage "iTask/modules/task/storage"
	"net/http"
)

//
//func ListItem(serviceCtx goservice.ServiceContext) func(ctx *gin.Context) {
//	return func(c *gin.Context) {
//		var queryString struct {
//			common.Paging
//			model.Filter
//		}
//
//		if err := c.ShouldBind(&queryString); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//			return
//		}
//
//		queryString.Paging.Process()
//
//		requester := c.MustGet(common.CurrentUser).(common.Requester)
//		db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)
//		apiItemCaller := serviceCtx.MustGet(common.PluginItemAPI).(interface {
//			GetServiceURL() string
//		})
//		store := storage.NewSQLStore(db)
//		likeStore := rpc.NewItemService(apiItemCaller.GetServiceURL(), serviceCtx.Logger("rpc.itemlikes"))
//		repo := repository.NewListItemRepo(store, likeStore, requester)
//		business := biz.NewListItemBiz(repo, requester)
//
//		result, err := business.ListItem(c.Request.Context(), &queryString.Filter, &queryString.Paging)
//
//		// SELECT * FROM todo_items ORDER BY id ASC LIMIT paging.Limit OFFSET (paging.Page - 1) * paging.Limit
//		if err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//			return
//		}
//
//		// fea_FakeID
//
//		for i := range result {
//			result[i].Mask()
//		}
//
//		c.JSON(http.StatusOK, common.NewSuccessResponse(result, queryString.Paging, queryString.Filter))
//	}
//}

func ListSimpleProjects(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var queryString struct {
			common.Paging
			model.Filter
		}

		// Default value for CreatedDayRange
		queryString.CreatedDayRange = -1

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		queryString.Paging.Process()

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		// !logging
		//fmt.Println("<-----> requester: ", requester)
		store := storage.NewSQLStore(db)
		taskStore := taskStorage.NewSQLStore(db)
		projectTagStore := projecTagStorage.NewSQLStore(db)
		projectMemberStore := storage2.NewSQLStore(db)
		repo := repository.NewListProjectRepo(store, taskStore, projectTagStore, projectMemberStore, requester)
		business := biz.NewListProjectBiz(repo, requester)

		result, err := business.ListSimpleProject(c.Request.Context(), &queryString.Paging)

		// SELECT * FROM todo_items ORDER BY id ASC LIMIT paging.Limit OFFSET (paging.Page - 1) * paging.Limit
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// fea_FakeID

		// for i := range result {
		// 	result[i].Mask()
		// }

		//for item := range result {
		//	result[item].PriorityValue = result[item].Priority.String()
		//	result[item].StatusValue = result[item].Status.String()
		//	result[item].PrivacyValue = result[item].Privacy.String()
		//}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, queryString.Paging, queryString.Filter))
	}
}
