package worker

import (
	"context"
	"iTask/common"
	"iTask/constant"
	"iTask/entities"
	"iTask/modules/project/model"
	"iTask/provider/mail"

	"github.com/go-redis/redis/v8"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const (
	QueueSendVerifyEmail       = "send_verify_email"
	QueueSendResetCodePassword = "send_verify_reset_code_password"
	QueueDefault               = "default"
)

// TaskProcessor will pick up the tasks from the Redis queue and process them

type TaskProcessor interface {
	Start() error
	ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
	ProcessTaskSendVerifyResetCodePassword(ctx context.Context, task *asynq.Task) error
	ProcessTaskSendConfirmBooking(ctx context.Context, task *asynq.Task) error
}

type AccountStorage interface {
	GetAccountByEmail(ctx context.Context, email string) (account *entities.Account, err error)
}

type BookingStorage interface {
	ListAllBookingWithCondition(ctx context.Context, condition []common.Condition) ([]entities.Booking, error)
}

type BookingUseCase interface {
	UpdateStatusBooking(ctx context.Context, bookingID, status int) error
}

type VerifyEmailsUseCase interface {
	CreateVerifyEmails(ctx context.Context, email string) (*entities.VerifyEmail, error)
	UpsertResetSetCodePassword(ctx context.Context, email string) (*entities.VerifyEmail, error)
	CreateProjectInvitationEmail(ctx context.Context, email string, project *model.Project) (*entities.VerifyEmail, error)
}

type redisTaskProcessor struct {
	server         *asynq.Server
	accountSto     AccountStorage
	verifyEmailsUC VerifyEmailsUseCase
	bookingSto     BookingStorage
	bookingUC      BookingUseCase
	mailer         mail.EmailSender
}

func NewRedisTaskProcessor(redisOpt *asynq.RedisClientOpt, accountSto AccountStorage, verifyEmailsUC VerifyEmailsUseCase, mailer mail.EmailSender, bookingSto BookingStorage, bookingUC BookingUseCase) TaskProcessor {

	logger := NewLogger()
	redis.SetLogger(logger)
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Queues: map[string]int{
				QueueSendVerifyEmail:       10,
				QueueSendResetCodePassword: 10,
				QueueDefault:               5,
			},
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				log.Error().Err(err).Str("task type", task.Type()).
					Bytes("payload", task.Payload()).
					Msg("error when process task")
			}),
			Logger: logger,
		})
	return &redisTaskProcessor{server: server, accountSto: accountSto, verifyEmailsUC: verifyEmailsUC, mailer: mailer, bookingSto: bookingSto, bookingUC: bookingUC}
}

func (processor *redisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(TaskSendConfirmBooking, processor.ProcessTaskSendConfirmBooking)
	mux.HandleFunc(TaskSendVerifyEmail, processor.ProcessTaskSendVerifyEmail)
	mux.HandleFunc(TaskSendInvitation, processor.ProcessTaskSendInvitation)
	mux.HandleFunc(TaskSendResetCodePassword, processor.ProcessTaskSendVerifyResetCodePassword)
	mux.HandleFunc(constant.TaskUpdateStatusBooking, processor.ProcessTaskUpdateStatusBooking)

	return processor.server.Start(mux)

}
