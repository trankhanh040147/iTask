package biz

import (
	"context"
	"iTask/common"
	model2 "iTask/modules/project/model"
	"iTask/modules/project_members/model"
)

type UpdateMemberStorage interface {
	GetMember(ctx context.Context, userId, projectId int) (*model.ProjectMember, error)
	UpdateMemberRole(ctx context.Context, projectId, memberId, role int) error
}

type updateMemberRoleBiz struct {
	store        UpdateMemberStorage
	projectStore ProjectStorage
}

func NewUpdateMemberRoleBiz(store UpdateMemberStorage, projectStore ProjectStorage) *updateMemberRoleBiz {
	return &updateMemberRoleBiz{store: store, projectStore: projectStore}
}

func (biz *updateMemberRoleBiz) UpdateMemberRole(ctx context.Context, projectId, memberId, role int) error {
	// todo: check permission

	// check project exist
	project, err := biz.projectStore.GetProject(ctx, map[string]interface{}{"id": projectId})
	if err != nil {
		return common.ErrCannotGetEntity(model.EntityName, err)
	} else if project == nil {
		return common.NewCustomError(nil, "project not found", "ErrProjectNotFound")
	}

	//log.Printf("projectId: %d, memberId: %d, role: %d", projectId, memberId, role)

	// check member exist
	_, err = biz.store.GetMember(ctx, memberId, projectId)
	if err != nil {
		return common.ErrCannotGetEntity(model.EntityName, err)
	}

	// check role is valid
	if role != model.RoleOwner && role != model.RoleMember {
		return common.NewCustomError(nil, "role is invalid", "ErrInvalidRole")
	}

	if err := biz.store.UpdateMemberRole(ctx, projectId, memberId, role); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	// todo: change old owner to member
	projectUpdate := model2.ProjectUpdate{CreatedBy: memberId}

	// update project
	if err = biz.projectStore.UpdateProject(ctx, map[string]interface{}{"id": projectId}, &projectUpdate); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	if err := biz.store.UpdateMemberRole(ctx, projectId, project.CreatedBy, model.RoleMember); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	return nil
}
