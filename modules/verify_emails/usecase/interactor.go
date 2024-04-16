package verifyemailsusecase

import (
	"context"
	"iTask/entities"
	"iTask/modules/project/model"
	projectMemberInvitedModel "iTask/modules/project_member_invited/model"
	projectMemberModel "iTask/modules/project_members/model"
)

type verifyEmailsStorage interface {
	Create(ctx context.Context, data *entities.VerifyEmail) (*entities.VerifyEmail, error)
	Get(ctx context.Context, email string, verifyCode string, _type int) (*entities.VerifyEmail, error)
	Update(ctx context.Context, email string, _type int, verifyEmail *entities.VerifyEmail) error
	UpdateEntity(ctx context.Context, cond map[string]interface{}, data *entities.VerifyEmail) error
	GetProjectInvitation(ctx context.Context, email string, verifyCode string, project_id int, _type int) (*entities.VerifyEmail, error)
	CheckIsExist(ctx context.Context, email string, _type int) (*entities.VerifyEmail, error)
}

type ProjectMemberInvitedStorage interface {
	CreateProjectMemberInvited(ctx context.Context, data *projectMemberInvitedModel.ProjectMemberInvited) error
}

type VerifyEmailsUseCase interface {
	CreateVerifyEmails(ctx context.Context, email string) (*entities.VerifyEmail, error)
	UpsertResetSetCodePassword(ctx context.Context, email string) (*entities.VerifyEmail, error)
	CreateProjectInvitationEmail(ctx context.Context, email string, project *model.Project) (*entities.VerifyEmail, error)
}
type AccountStorage interface {
	UpdateIsVerifyEmailByEmail(ctx context.Context, email string) error
	GetAccountByEmail(ctx context.Context, email string) (account *entities.Account, err error)
}

type ProjectMemberStorage interface {
	CreateProjectMember(ctx context.Context, data *projectMemberModel.ProjectMemberCreation) error
	GetMember(ctx context.Context, userId, projectId int) (*projectMemberModel.ProjectMember, error)
}

type verifyEmailsUseCase struct {
	verifyEmailsStore           verifyEmailsStorage
	projectMemberInvitedStorage ProjectMemberInvitedStorage
	accountStore                AccountStorage
	projectMemberStore          ProjectMemberStorage
}

func NewVerifyEmailsUseCase(verifyEmailsStore verifyEmailsStorage, accountSto AccountStorage, projectMemberInvitedStorage ProjectMemberInvitedStorage, projectMemberStore ProjectMemberStorage) *verifyEmailsUseCase {
	return &verifyEmailsUseCase{
		verifyEmailsStore:           verifyEmailsStore,
		accountStore:                accountSto,
		projectMemberInvitedStorage: projectMemberInvitedStorage,
		projectMemberStore:          projectMemberStore,
	}
}
