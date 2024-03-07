package repository

import (
	"context"
	"iTask/common"
	"iTask/modules/task/model"
)

type ListTaskStorage interface {
	ListTask(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.Task, error)
}

//type TaskStorage interface {
//	GetTotalTasks(ctx context.Context, cond map[string]interface{}) (map[int]int, error)
//}

type listTaskRepo struct {
	store ListTaskStorage
	//taskStorage TaskStorage
	requester common.Requester
}

func NewListTaskRepo(store ListTaskStorage, requester common.Requester) *listTaskRepo {
	return &listTaskRepo{store: store, requester: requester}
}

func (repo *listTaskRepo) ListTask(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.Task, error) {
	newCtx := context.WithValue(ctx, common.CurrentUser, repo.requester)

	data, err := repo.store.ListTask(newCtx, filter, paging, moreKeys...)
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

	//TotalTasksMap, err := repo.taskStorage.GetTotalTasks(newCtx, nil)
	//if err != nil {
	//	return data, nil
	//}

	//cond := map[string]interface{}{"status": common.StatusCompleted}

	//TotalCompletedTasksMap, err := repo.taskStorage.GetTotalTasks(newCtx, cond)
	//if err != nil {
	//	return data, nil
	//}
	//
	//for index := range data {
	//	data[index].TotalTasks = TotalTasksMap[data[index].Id]
	//	data[index].TotalCompletedTasks = TotalCompletedTasksMap[data[index].Id]
	//}

	return data, nil
}
