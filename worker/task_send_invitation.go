package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"iTask/constant"
	"iTask/entities"
	"iTask/modules/project/model"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const (
	TaskSendInvitation = "task:send_project_invitation"
	UrlInvitation      = constant.URL_HOST_EC2 + "/project/invitation"
)

type PayLoadSendInvitationEmail struct {
	Email   string         `json:"email"`
	Project *model.Project `json:"project"`
}

func (distributor *redisTaskDistributor) DistributeTaskSendInvitationEmail(
	ctx context.Context,
	payload *PayLoadSendInvitationEmail,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error when marshal payload: %v", err)
	}

	task := asynq.NewTask(TaskSendInvitation, jsonPayload, opts...)

	// send task to Redis queue
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("error when enqueue task: %v", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task")

	return nil
}

func (processor *redisTaskProcessor) ProcessTaskSendInvitation(ctx context.Context, task *asynq.Task) error {

	// log start processing task
	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).Msg("processing task")

	var payload PayLoadSendInvitationEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("error when unmarshal payload: %w", asynq.SkipRetry)
	}

	account, err := processor.accountSto.GetAccountByEmail(ctx, payload.Email)

	if err != nil {
		return fmt.Errorf("error when get account by email: %w", err)
	}

	dataVerifyEmail, err := processor.verifyEmailsUC.CreateProjectInvitationEmail(ctx, account.Email, payload.Project)
	if err != nil {
		return fmt.Errorf("error when create verify email: %w", err)
	}

	if err := sendInvitationEmail(processor, dataVerifyEmail, account, payload.Project); err != nil {
		return fmt.Errorf("error when send invitation: %w", err)
	}

	log.Info().Msg("send invitation success")

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("email", account.Email).Msg("processed task")
	return nil
}

func sendInvitationEmail(processor *redisTaskProcessor, verifyEmail *entities.VerifyEmail, account *entities.Account, project *model.Project) error {
	subject := "Welcome to ITask"
	verifyUrl := fmt.Sprintf("%s?email=%s&secret_code=%s&project_id=%d",
		UrlInvitation, verifyEmail.Email, verifyEmail.ScretCode, project.Id)
	content := fmt.Sprintf(`Hello %s,<br/>
	You have been invited to join project %s !<br/>
	Please <a href="%s">click here</a> to join the project.<br/>
	`, account.FullName, project.Name, verifyUrl)
	to := []string{account.Email}

	err := processor.mailer.SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to send verify email: %w", err)
	}
	return nil
}
