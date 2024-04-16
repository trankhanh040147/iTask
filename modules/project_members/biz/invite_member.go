package biz

import (
	"context"
	"github.com/hibiken/asynq"
	"iTask/common"
	"iTask/entities"
	projectModel "iTask/modules/project/model"
	"iTask/modules/project_members/model"
	projectMemberStore "iTask/modules/project_members/storage"
	"iTask/worker"
	"time"
)

type InviteMemberStore interface {
	GetMember(ctx context.Context, userId, projectId int) (*model.ProjectMember, error)
	InviteMemberTx(ctx context.Context, createUserTxParam projectMemberStore.InviteMemberTxParam) error
}

type AccountStorage interface {
	GetAccountByEmail(ctx context.Context, email string) (*entities.Account, error)
}

type ProjectStorage interface {
	GetProject(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*projectModel.Project, error)
	UpdateProject(ctx context.Context, cond map[string]interface{}, dataUpdate *projectModel.ProjectUpdate) error
}

type inviteMemberBiz struct {
	store           InviteMemberStore
	accountStore    AccountStorage
	projectStore    ProjectStorage
	taskDistributor worker.TaskDistributor
}

func NewInviteMemberBiz(store InviteMemberStore, accountStore AccountStorage, projectStore ProjectStorage, taskDistributor worker.TaskDistributor) *inviteMemberBiz {
	return &inviteMemberBiz{
		store:           store,
		accountStore:    accountStore,
		projectStore:    projectStore,
		taskDistributor: taskDistributor,
	}
}

func (biz *inviteMemberBiz) InviteMember(ctx context.Context, email string, projectId int) error {
	// check if account with email exists
	account, err := biz.accountStore.GetAccountByEmail(ctx, email)
	if account == nil {
		return common.ErrEmailNotExist("Account", err)
	}

	// check if account with email is already a member of the project
	userId := account.Id
	member, err := biz.store.GetMember(ctx, userId, projectId)
	if member != nil {
		return common.ErrEntityExisted(model.EntityName)
	}

	// todo: get project by projectId, check if project is existed
	project, err := biz.projectStore.GetProject(ctx, map[string]interface{}{"id": projectId})
	if err != nil {
		return err
	}
	if project == nil {
		return common.ErrEntityNotFound(projectModel.EntityName)
	}

	// send email to invite member to join project

	paramInviteTx := projectMemberStore.InviteMemberTxParam{
		Data: account,
		// todo: add project to payload
		AfterCreate: func(data *entities.Account) error {
			// after other logics success, we will send email for user to accept invitation
			taskPayload := worker.PayLoadSendInvitationEmail{
				Email:   account.Email,
				Project: project,
			}
			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(10 * time.Second),
				asynq.Queue(worker.QueueSendVerifyEmail),
			}

			return biz.taskDistributor.DistributeTaskSendInvitationEmail(ctx, &taskPayload, opts...)
		},
	}

	//if err = uc.accountStorage.CreateTx(ctx, paramInviteTx); err != nil {
	if err = biz.store.InviteMemberTx(ctx, paramInviteTx); err != nil {
		return err
	}

	return nil
}
