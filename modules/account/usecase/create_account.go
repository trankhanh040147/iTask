package accountusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/constant"
	"paradise-booking/entities"
	"paradise-booking/modules/account/convert"
	"paradise-booking/modules/account/iomodel"
	accountstorage "paradise-booking/modules/account/storage"
	"paradise-booking/utils"
	"paradise-booking/worker"
	"time"

	"github.com/hibiken/asynq"
)

func (uc *accountUseCase) CreateAccount(ctx context.Context, accountModel *iomodel.AccountRegister) (result *string, err error) {

	// convert from iomodel to entity
	accountEntity := convert.ConvertAccountRegisModelToEntity(accountModel)
	accountEntity.Status = int(constant.StatusActive)
	accountEntity.IsEmailVerified = 1 // change_later
	// hash password before store in db
	hashedPassword, err := utils.HashPassword(accountEntity.Password)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	// default in first register account will have role user
	accountEntity.Role = int(constant.UserRole)
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
