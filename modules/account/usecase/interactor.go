package accountusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/config"
	"paradise-booking/entities"
	accountstorage "paradise-booking/modules/account/storage"
	"paradise-booking/worker"
)

type AccountStorage interface {
	Create(ctx context.Context, account *entities.Account) (err error)
	GetAccountByEmail(ctx context.Context, email string) (account *entities.Account, err error)
	UpdateAccountById(ctx context.Context, id int, accountUpdate *entities.Account) error
	GetProfileByID(ctx context.Context, id int) (*entities.Account, error)
	CreateTx(ctx context.Context, createUserTxParam accountstorage.CreateUserTxParam) error
	GetAllAccountUserAndVendor(ctx context.Context, paging *common.Paging) ([]entities.Account, error)
}

type VerifyEmailsUseCase interface {
	UpsertResetSetCodePassword(ctx context.Context, email string) (*entities.VerifyEmail, error)
}

type accountUseCase struct {
	accountStorage  AccountStorage
	verifyEmailsUC  VerifyEmailsUseCase
	cfg             *config.Config
	taskDistributor worker.TaskDistributor
}

func NewUserUseCase(cfg *config.Config, accountSto AccountStorage, verifyEmailsUC VerifyEmailsUseCase, taskDistributor worker.TaskDistributor) *accountUseCase {
	return &accountUseCase{accountSto, verifyEmailsUC, cfg, taskDistributor}
}
