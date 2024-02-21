package biz

import (
	"context"
	"iTask/common"
	"iTask/modules/project/model"
)

type ListProjectRepo interface {
	ListProject(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.Project, error)
}

type listProjectBiz struct {
	repo      ListProjectRepo
	requester common.Requester
}

func NewListProjectBiz(repo ListProjectRepo, requester common.Requester) *listProjectBiz {
	return &listProjectBiz{repo: repo, requester: requester}
}

func (biz *listProjectBiz) ListProject(ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
) ([]model.Project, error) {
	newCtx := context.WithValue(ctx, common.CurrentUser, biz.requester)

	data, err := biz.repo.ListProject(newCtx, filter, paging, "Owner")

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return data, nil
}
