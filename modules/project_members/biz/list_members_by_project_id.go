package biz

import (
	"context"
	"iTask/common"
	"iTask/modules/project_members/model"
)

type ListMembersByIdStore interface {
	ListMembersById(
		ctx context.Context,
		id int,
		moreKeys ...string,
	) ([]model.SimpleMember, error)
}

type listProjectMemberByIdBiz struct {
	store ListMembersByIdStore
}

func NewListMembersByIdBiz(store ListMembersByIdStore) *listProjectMemberByIdBiz {
	return &listProjectMemberByIdBiz{store: store}
}

func (biz *listProjectMemberByIdBiz) ListMembersById(ctx context.Context, id int) ([]model.SimpleMember, error) {

	// get project_ids from project_members where user_id = requester.GetUserId()

	data, err := biz.store.ListMembersById(ctx, id, "Owner")

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return data, nil
}
