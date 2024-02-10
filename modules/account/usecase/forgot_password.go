package accountusecase

import (
	"context"
	"errors"
	"paradise-booking/common"
	"paradise-booking/worker"
	"time"

	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

func (uc *accountUseCase) ForgotPassword(ctx context.Context, email string) error {
	// check email exist
	_, err := uc.accountStorage.GetAccountByEmail(ctx, email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrEmailNotExist("account", errors.New("email not exist"))
		}
		return err
	}

	// send email

	taskPayload := worker.PayloadSendVerifyResetCodePassword{
		Email: email,
	}
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
		asynq.Queue(worker.QueueSendResetCodePassword),
	}
	_ = uc.taskDistributor.DistributeTaskSendVerifyResetCodePassword(ctx, &taskPayload, opts...)
	return nil
}
