package worker

import (
	"context"

	"github.com/hibiken/asynq"
)

type TaskDistributor interface {
	DistributeTaskSendVerifyEmail(
		ctx context.Context,
		payload *PayloadSendVerifyEmail,
		opts ...asynq.Option) error
	DistributeTaskSendInvitationEmail(
		ctx context.Context,
		payload *PayLoadSendInvitationEmail,
		opts ...asynq.Option,
	) error
	DistributeTaskSendVerifyResetCodePassword(
		ctx context.Context,
		payload *PayloadSendVerifyResetCodePassword,
		opts ...asynq.Option,
	) error
	DistributeTaskSendConfirmBooking(
		ctx context.Context,
		payload *PayloadSendConfirmBooking,
		opts ...asynq.Option,
	) error
	DistributeTaskUpdateStatusBooking(ctx context.Context, opts ...asynq.Option) error
}

type redisTaskDistributor struct {
	client *asynq.Client
}

func NewRedisTaskDistributor(redisOpt *asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(redisOpt)
	return &redisTaskDistributor{client: client}
}
