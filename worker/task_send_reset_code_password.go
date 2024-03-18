package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"iTask/entities"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const (
	TaskSendResetCodePassword = "task:send_verify_reset_code_password"
	//UrlVerifyEmail      = "https://paradisebookingapp.up.railway.app/api/v1/verify_email"
)

type PayloadSendVerifyResetCodePassword struct {
	Email string `json:"email"`
}

func (distributor *redisTaskDistributor) DistributeTaskSendVerifyResetCodePassword(
	ctx context.Context,
	payload *PayloadSendVerifyResetCodePassword,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error when marshal payload: %v", err)
	}
	task := asynq.NewTask(TaskSendResetCodePassword, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("error when enqueue task: %v", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task")
	return nil
}

func (processor *redisTaskProcessor) ProcessTaskSendVerifyResetCodePassword(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendVerifyResetCodePassword
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("error when unmarshal payload: %w", asynq.SkipRetry)
	}

	account, err := processor.accountSto.GetAccountByEmail(ctx, payload.Email)
	// if err == gorm.ErrRecordNotFound {
	// 	return fmt.Errorf("account with email %s not found: %w", payload.Email, asynq.SkipRetry)
	// }
	if err != nil {
		return fmt.Errorf("error when get account by email: %w", err)
	}

	dataResetCodePass, err := processor.verifyEmailsUC.UpsertResetSetCodePassword(ctx, account.Email)
	if err != nil {
		return fmt.Errorf("error when create verify email: %w", err)
	}

	sendMailToVerifyResetCodePassword(processor, dataResetCodePass, account)
	log.Info().Msg("send verify code password success")

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("email", account.Email).Msg("processed task")
	return nil
}

func sendMailToVerifyResetCodePassword(processor *redisTaskProcessor, verifyResetCodePassword *entities.VerifyEmail, account *entities.Account) error {
	subject := "Welcome to ITask"
	content := fmt.Sprintf(`Hello %s,<br/>
	This is reset code for you: %s<br/>
	`, account.FullName, verifyResetCodePassword.ScretCode)
	to := []string{account.Email}

	err := processor.mailer.SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to send verify reset code password: %w", err)
	}
	return nil
}
