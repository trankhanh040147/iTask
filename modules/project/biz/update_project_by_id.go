package biz

import (
	"context"
	"errors"
	"iTask/common"
	"iTask/modules/project/model"
)

type UpdateProjectStorage interface {
	GetProject(ctx context.Context, cond map[string]interface{}) (*model.Project, error)
	UpdateProject(ctx context.Context, cond map[string]interface{}, dataUpdate *model.ProjectUpdate) error
}

type updateProjectBiz struct {
	store     UpdateProjectStorage
	requester common.Requester
}

func NewUpdateProjectBiz(store UpdateProjectStorage, requester common.Requester) *updateProjectBiz {
	return &updateProjectBiz{store: store, requester: requester}
}

func (biz *updateProjectBiz) UpdateProject(ctx context.Context, id int, dataUpdate *model.ProjectUpdate) error {
	data, err := biz.store.GetProject(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if data.Status.String() == "Deleted" {
		return common.ErrEntityDeleted(model.EntityName)
	}

	isOwner := data.CreatedBy == biz.requester.GetUserId()

	if !isOwner && !common.IsAdmin(biz.requester) {
		return common.ErrNoPermission(errors.New("you don't have permission to update this item"))
	}

	if err := biz.store.UpdateProject(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	return nil
}
