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
	GetTotalTasks(ctx context.Context, ids []int) (map[int]int, error)
}

type listProjectRepo struct {
	store       ListProjectStorage
	taskStorage TaskStorage
	requester   common.Requester
}

//	func NewListProjectRepo(store ListProjectStorage, taskStorage TaskStorage, requester common.Requester) *listProjectRepo {
//		return &listProjectRepo{store: store, taskStorage: taskStorage, requester: requester}
//	}

func NewListProjectRepo(store ListProjectStorage, requester common.Requester) *listProjectRepo {
	return &listProjectRepo{store: store, requester: requester}
}

func (repo *listProjectRepo) ListProject(
	ctx context.Context,
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

	//ids := make([]int, len(data))
	//for index := range ids {
	//	ids[index] = data[index].Id
	//}
	//
	//ProjectLikesMap, err := repo.taskStorage.GetTotalTasks(newCtx, ids)
	//if err != nil {
	//	return data, nil
	//}
	//
	//for index := range data {
	//	data[index].TotalTasks = ProjectLikesMap[data[index].Id]
	//}

	return data, nil
}
