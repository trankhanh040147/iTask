package biz

import (
	"context"
	"iTask/common"
	"iTask/modules/project_members/model"
)

type DeleteMemberStorage interface {
	GetMember(ctx context.Context, userId, projectId int) (*model.ProjectMember, error)
	DeleteMember(ctx context.Context, cond map[string]interface{}) error
}

type deleteMemberBiz struct {
	store DeleteMemberStorage
}

func NewDeleteMemberBiz(store DeleteMemberStorage) *deleteMemberBiz {
	return &deleteMemberBiz{store: store}
}

func (biz *deleteMemberBiz) DeleteMember(ctx context.Context, userId, projectId int) error {
	_, err := biz.store.GetMember(ctx, userId, projectId)

	if err != nil {
		return common.ErrCannotGetEntity(model.EntityName, err)
	}

	if err := biz.store.DeleteMember(ctx, map[string]interface{}{"user_id": userId, "project_id": projectId}); err != nil {
		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}

	return nil
}
