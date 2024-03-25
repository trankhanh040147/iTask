package biz

import (
	"context"
	"iTask/common"
	"iTask/modules/task/model"
)

type ListTaskRepo interface {
	ListTask(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([](*model.Task), error)
}

type listTaskBiz struct {
	repo      ListTaskRepo
	requester common.Requester
}

func NewListTaskBiz(repo ListTaskRepo, requester common.Requester) *listTaskBiz {
	return &listTaskBiz{repo: repo, requester: requester}
}

func (biz *listTaskBiz) ListTask(ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
) ([](*model.Task), error) {
	newCtx := context.WithValue(ctx, common.CurrentUser, biz.requester)

	data, err := biz.repo.ListTask(newCtx, filter, paging, "Owner")

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	for _, t := range data {
		t.Parsing()
		//fmt.Printf("Task: %+v\n", t)
	}

	return data, nil
}
