package repository

import (
	"context"
	"iTask/common"
	"iTask/modules/project/model"
)

type ListProjectStorage interface {
	ListProject(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.Project, error)
}

type TaskStorage interface {
	GetTotalTasks(ctx context.Context, cond map[string]interface{}) (map[int]int, error)
	//ListTask(ctx context.Context, filter *model.Filter, paging *common.Paging, moreKeys ...string) ([]model.Task, error)
}

type ProjectTagStorage interface {
	GetProjectTagsByProjectId(ctx context.Context, cond map[string]interface{}) (map[int]string, error)
}

type listProjectRepo struct {
	store             ListProjectStorage
	taskStorage       TaskStorage
	ProjectTagStorage ProjectTagStorage
	requester         common.Requester
}

func NewListProjectRepo(store ListProjectStorage, taskStorage TaskStorage, ProjectTagStorage ProjectTagStorage, requester common.Requester) *listProjectRepo {
	return &listProjectRepo{store: store, taskStorage: taskStorage, ProjectTagStorage: ProjectTagStorage, requester: requester}
}

//func NewListProjectRepo(store ListProjectStorage, requester common.Requester) *listProjectRepo {
//	return &listProjectRepo{store: store, requester: requester}
//}

func (repo *listProjectRepo) ListProject(ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.Project, error) {
	newCtx := context.WithValue(ctx, common.CurrentUser, repo.requester)

	data, err := repo.store.ListProject(newCtx, filter, paging, moreKeys...)
	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	if len(data) == 0 {
		return data, nil
	}

	ids := make([]int, len(data))
	for index := range ids {
		ids[index] = data[index].Id
	}

	// get total tasks and completed tasks
	TotalTasksMap, err := repo.taskStorage.GetTotalTasks(newCtx, nil)
	if err != nil {
		return data, nil
	}

	cond := map[string]interface{}{"status": common.StatusCompleted}

	TotalCompletedTasksMap, err := repo.taskStorage.GetTotalTasks(newCtx, cond)
	if err != nil {
		return data, nil
	}

	for index := range data {
		data[index].TotalTasks = TotalTasksMap[data[index].Id]
		data[index].TotalCompletedTasks = TotalCompletedTasksMap[data[index].Id]
	}

	// get tags
	ProjectTagsMap, err := repo.ProjectTagStorage.GetProjectTagsByProjectId(newCtx, nil)
	if err != nil {
		return data, nil
	}

	cond = map[string]interface{}{"status": common.StatusCompleted}

	for index := range data {
		data[index].Tags = ProjectTagsMap[data[index].Id]
	}

	return data, nil
}
