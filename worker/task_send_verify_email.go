package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"iTask/constant"
	"iTask/entities"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const (
	TaskSendVerifyEmail = "task:send_verify_email"
	UrlVerifyEmail      = constant.URL_HOST_EC2 + "/verify_email"
)

type PayloadSendVerifyEmail struct {
	Email string `json:"email"`
}

func (distributor *redisTaskDistributor) DistributeTaskSendVerifyEmail(
	ctx context.Context,
	payload *PayloadSendVerifyEmail,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error when marshal payload: %v", err)
	}

	task := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opts...)

	// send task to Redis queue
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("error when enqueue task: %v", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task")

	return nil
}

func (processor *redisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {

	// log start processing task
	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).Msg("processing task")

	var payload PayloadSendVerifyEmail
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

	dataVerifyEmail, err := processor.verifyEmailsUC.CreateVerifyEmails(ctx, account.Email)
	if err != nil {
		return fmt.Errorf("error when create verify email: %w", err)
	}

	if err := sendMailToVerifyEmail(processor, dataVerifyEmail, account); err != nil {
		return fmt.Errorf("error when send verify email: %w", err)
	}

	log.Info().Msg("send verify email success")

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("email", account.Email).Msg("processed task")
	return nil
}

func sendMailToVerifyEmail(processor *redisTaskProcessor, verifyEmail *entities.VerifyEmail, account *entities.Account) error {
	subject := "Welcome to ITask"
	verifyUrl := fmt.Sprintf("%s?email=%s&secret_code=%s",
		UrlVerifyEmail, verifyEmail.Email, verifyEmail.ScretCode)
	content := fmt.Sprintf(`Hello %s,<br/>
	Thank you for registering with us!<br/>
	Please <a href="%s">click here</a> to verify your email address.<br/>
	`, account.FullName, verifyUrl)
	to := []string{account.Email}

	err := processor.mailer.SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to send verify email: %w", err)
	}
	return nil
}
