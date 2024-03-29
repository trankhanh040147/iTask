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
	ListSimpleProjects(ctx context.Context, paging *common.Paging, moreKeys ...string,
	) ([]model.SimpleProject, error)
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

	// get project_ids from project_members where user_id = requester.GetUserId()

	data, err := biz.repo.ListProject(newCtx, filter, paging, "Owner")

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return data, nil
}

func (biz *listProjectBiz) ListSimpleProject(ctx context.Context, paging *common.Paging) ([]model.SimpleProject, error) {
	newCtx := context.WithValue(ctx, common.CurrentUser, biz.requester)

	data, err := biz.repo.ListSimpleProjects(newCtx, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return data, nil
}
