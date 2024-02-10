package accounthandler

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/config"
	"paradise-booking/entities"
	"paradise-booking/modules/account/iomodel"
	jwtprovider "paradise-booking/provider/jwt"
)

type accountUseCase interface {
	CreateAccount(ctx context.Context, accountModel *iomodel.AccountRegister) (result *string, err error)
	LoginAccount(ctx context.Context, accountModel *iomodel.AccountLogin) (toke *jwtprovider.Token, err error)
	UpdatePersonalInforAccountById(ctx context.Context, accountModel *iomodel.AccountUpdatePersonalInfo, id int) (err error)
	GetAccountByEmail(ctx context.Context, email string) (account *iomodel.AccountInfoResp, err error)
	GetAccountByID(ctx context.Context, id int) (account *iomodel.AccountInfoResp, err error)
	UpdateAccountRoleByID(ctx context.Context, accountModel *iomodel.AccountChangeRole, id int) (err error)
	GetAllAccountUserAndVendor(ctx context.Context, paging *common.Paging) ([]entities.Account, error)
	ChangePassword(ctx context.Context, email string, changePassModel *iomodel.ChangePassword) error
	ChangeStatusAccount(ctx context.Context, accountID int, status int) error
	ForgotPassword(ctx context.Context, email string) error
	UpdatePassword(ctx context.Context, email string, newPassword string) error
}

type accountHandler struct {
	accountUC accountUseCase
	cfg       *config.Config
}

func NewAccountHandler(cfg *config.Config, accountUseCase accountUseCase) *accountHandler {
	return &accountHandler{accountUC: accountUseCase, cfg: cfg}
}
