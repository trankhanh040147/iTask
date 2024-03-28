package verifyemailsusecase

import (
	"context"
	"iTask/entities"
	"iTask/modules/project/model"
)

type verifyEmailsStorage interface {
	Create(ctx context.Context, data *entities.VerifyEmail) (*entities.VerifyEmail, error)
	Get(ctx context.Context, email string, verifyCode string, _type int) (*entities.VerifyEmail, error)
	CheckIsExist(ctx context.Context, email string, _type int) (*entities.VerifyEmail, error)
	Update(ctx context.Context, email string, _type int, verifyEmail *entities.VerifyEmail) error
}

type VerifyEmailsUseCase interface {
	CreateVerifyEmails(ctx context.Context, email string) (*entities.VerifyEmail, error)
	UpsertResetSetCodePassword(ctx context.Context, email string) (*entities.VerifyEmail, error)
	CreateProjectInvitationEmail(ctx context.Context, email string, project *model.Project) (*entities.VerifyEmail, error)
}

type AccountStorage interface {
	UpdateIsVerifyEmailByEmail(ctx context.Context, email string) error
}

type verifyEmailsUseCase struct {
	verifyEmailsStore verifyEmailsStorage
	accountStore      AccountStorage
}

func NewVerifyEmailsUseCase(verifyEmailsStore verifyEmailsStorage, accountSto AccountStorage) *verifyEmailsUseCase {
	return &verifyEmailsUseCase{verifyEmailsStore: verifyEmailsStore, accountStore: accountSto}
}
