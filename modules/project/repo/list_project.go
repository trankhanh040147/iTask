package repository

import (
	"context"
	"fmt"
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
	ListProjectByProjectIds(
		ctx context.Context,
		projectIds []int,
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

type ProjectMemberStorage interface {
	GetProjectIdsByUserId(ctx context.Context, userId int) ([]int, error)
}

type listProjectRepo struct {
	store                ListProjectStorage
	taskStorage          TaskStorage
	projectTagStorage    ProjectTagStorage
	projectMemberStorage ProjectMemberStorage
	requester            common.Requester
}

func NewListProjectRepo(store ListProjectStorage, taskStorage TaskStorage, ProjectTagStorage ProjectTagStorage, ProjectMemberStorage ProjectMemberStorage, requester common.Requester) *listProjectRepo {
	return &listProjectRepo{store: store, taskStorage: taskStorage, projectTagStorage: ProjectTagStorage, projectMemberStorage: ProjectMemberStorage, requester: requester}
}

func (repo *listProjectRepo) ListProject(ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.Project, error) {
	newCtx := context.WithValue(ctx, common.CurrentUser, repo.requester)

	// get project_ids of projects where Requester is a member
	projectIds, err := repo.projectMemberStorage.GetProjectIdsByUserId(newCtx, repo.requester.GetUserId())

	// !logging
	fmt.Println("<-----> projectIds: ", projectIds)

	//data, err := repo.store.ListProject(newCtx, filter, paging, moreKeys...)
	data, err := repo.store.ListProjectByProjectIds(newCtx, projectIds, filter, paging, moreKeys...)
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
	ProjectTagsMap, err := repo.projectTagStorage.GetProjectTagsByProjectId(newCtx, nil)
	if err != nil {
		return data, nil
	}

	for index := range data {
		data[index].Tags = ProjectTagsMap[data[index].Id]
	}

	return data, nil
}
