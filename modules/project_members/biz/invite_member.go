package biz

import (
	"context"
	"iTask/common"
	"iTask/entities"
	"iTask/modules/project_members/model"
)

type GetMemberStore interface {
	GetMember(ctx context.Context, userId, projectId int) (*model.ProjectMember, error)
}

type AccountStorage interface {
	GetAccountByEmail(ctx context.Context, email string) (*entities.Account, error)
}

type inviteMemberBiz struct {
	store        GetMemberStore
	accountStore AccountStorage
}

func NewInviteMemberBiz(store GetMemberStore, accountStore AccountStorage) *inviteMemberBiz {
	return &inviteMemberBiz{store: store, accountStore: accountStore}
}

func (biz *inviteMemberBiz) InviteMember(ctx context.Context, email string, projectId int) error {
	// check if account with email exists
	accountFound, err := biz.accountStore.GetAccountByEmail(ctx, email)
	if accountFound == nil {
		return common.ErrEmailNotExist("Account", err)
	}

	// check if account with email is already a member of the project
	userId := accountFound.Id
	member, err := biz.store.GetMember(ctx, userId, projectId)
	if member != nil {
		return common.ErrEntityExisted("Project member")
	}

	// send email to invite member to join project
	// todo: create task send mail (w/ transaction)
	return nil
}
