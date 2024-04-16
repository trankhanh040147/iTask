package biz

import (
	"context"
	"iTask/common"
	"iTask/entities"
	projectModel "iTask/modules/project/model"
	"iTask/modules/project_members/model"
	"time"
)

type FindUninvitedMemberStore interface {
	GetMember(ctx context.Context, userId, projectId int) (*model.ProjectMember, error)
}

type EmailStorage interface {
	FindProjectInvitation(ctx context.Context, cond map[string]interface{}) ([]entities.VerifyEmail, error)
}
type findUninvitedMemberBiz struct {
	store        FindUninvitedMemberStore
	accountStore AccountStorage
	projectStore ProjectStorage
	emailStorage EmailStorage
}

func NewFindUninvitedMemberBiz(store FindUninvitedMemberStore, accountStore AccountStorage, projectStore ProjectStorage, emailStorage EmailStorage) *findUninvitedMemberBiz {
	return &findUninvitedMemberBiz{
		store:        store,
		accountStore: accountStore,
		projectStore: projectStore,
		emailStorage: emailStorage,
	}
}

func (biz *findUninvitedMemberBiz) FindUninvitedMember(ctx context.Context, projectId int, userEmail string) (*entities.Account, error) {
	// *check if account with email exists
	account, err := biz.accountStore.GetAccountByEmail(ctx, userEmail)
	if err != nil {
		return nil, common.ErrCannotGetEntity(entities.EntityName, err)
	}

	// *check if account with email is already a member of the project
	userId := account.Id
	member, err := biz.store.GetMember(ctx, userId, projectId)
	if member != nil {
		return nil, common.ErrEntityExisted(model.EntityName)
	}

	// *get project by projectId, check if project is existed
	project, err := biz.projectStore.GetProject(ctx, map[string]interface{}{"id": projectId})
	if err != nil {
		return nil, err
	}
	if project == nil {
		return nil, common.ErrEntityNotFound(projectModel.EntityName)
	}

	// set account invited status to false
	account.IsInvited = false

	// *check if account has been invited and the invitation is still valid
	invitations, err := biz.emailStorage.FindProjectInvitation(ctx, map[string]interface{}{
		"project_id": projectId, "email": userEmail, "is_used": 0, "type": common.VerifyEmailTypeProjectInvitation,
	})
	var invitation entities.VerifyEmail
	if err != nil {
		return nil, common.ErrCannotGetEntity("verify_emails", err)
	}
	if len(invitations) == 0 {
		return account, nil
	} else {
		invitation = invitations[0]
	}

	// todo: the invitation is not expired yet
	currentTime := time.Now()
	if !invitation.ExpiredAt.Before(currentTime) {
		account.IsInvited = true
	}

	return account, nil
}
