package accountusecase

import (
	"context"
	"errors"
	"iTask/common"
	"iTask/entities"
	"iTask/modules/account/convert"
	"iTask/modules/account/iomodel"
	accountstorage "iTask/modules/account/storage"
	"iTask/utils"
	"iTask/worker"
	"time"

	"github.com/hibiken/asynq"
)

func (uc *accountUseCase) CreateAccount(ctx context.Context, accountModel *iomodel.AccountRegister) (result *string, err error) {

	// convert from iomodel to entity
	accountEntity := convert.ConvertAccountRegisModelToEntity(accountModel)

	// check if email is existed
	accountFound, err := uc.accountStorage.GetAccountByEmail(ctx, accountModel.Email)

	if accountFound != nil {
		return nil, errors.New("Email is existed")
	}

	accountEntity.Status = entities.StatusActive
	accountEntity.IsEmailVerified = 1 // change_later
	// hash password before store in db
	hashedPassword, err := utils.HashPassword(accountEntity.Password)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	// default in first register account will have role user
	accountEntity.Role = entities.RoleMember
	accountEntity.Password = hashedPassword

	paramCreateTx := accountstorage.CreateUserTxParam{
		Data: &accountEntity,
		AfterCreate: func(data *entities.Account) error {
			// after create account success, we will send email to user to verify account
			taskPayload := worker.PayloadSendVerifyEmail{
				Email: accountEntity.Email,
			}
			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(10 * time.Second),
				asynq.Queue(worker.QueueSendVerifyEmail),
			}

			return uc.taskDistributor.DistributeTaskSendVerifyEmail(ctx, &taskPayload, opts...)
		},
	}

	if err = uc.accountStorage.CreateTx(ctx, paramCreateTx); err != nil {
		return nil, err
	}

	return &accountEntity.Email, nil
}
