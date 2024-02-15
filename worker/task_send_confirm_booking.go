package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"paradise-booking/constant"
	"paradise-booking/entities"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const (
	TaskSendConfirmBooking = "task:send_confirm_booking"
	UrlConfirmBooking      = constant.URL_HOST_EC2 + "/confirm_booking"
)

type PayloadSendConfirmBooking struct {
	Email     string `json:"email"`
	BookingID int    `json:"booking_id"`
}

func (distributor *redisTaskDistributor) DistributeTaskSendConfirmBooking(
	ctx context.Context,
	payload *PayloadSendConfirmBooking,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error when marshal payload: %v", err)
	}
	task := asynq.NewTask(TaskSendConfirmBooking, jsonPayload, opts...)

	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("error when enqueue task: %v", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task")
	return nil
}

func (processor *redisTaskProcessor) ProcessTaskSendConfirmBooking(ctx context.Context, task *asynq.Task) error {
	log.Info().Msg("process task send confirm booking")
	var payload PayloadSendConfirmBooking
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

	sendMailToVerifyBooking(processor, account, payload.BookingID)
	log.Info().Msg("send verify booking success")

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("email", account.Email).Msg("processed task")

	return nil
}

func sendMailToVerifyBooking(processor *redisTaskProcessor, account *entities.Account, bookingID int) error {
	subject := "Welcome to Paradise Booking"
	verifyUrl := fmt.Sprintf("%s?booking_id=%d&status=%d",
		UrlConfirmBooking, bookingID, constant.BookingStatusConfirmed)
	content := fmt.Sprintf(`Hello %s,<br/>
	Thank you for booking with us!<br/>
	Please <a href="%s">click here</a> to confirm your booking.<br/>
	`, account.FullName, verifyUrl)
	to := []string{account.Email}

	err := processor.mailer.SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to send confirm booking: %w", err)
	}
	return nil
}
