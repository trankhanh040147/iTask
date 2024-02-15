package worker

import (
	"context"
	"iTask/common"
	"iTask/config"
	"iTask/constant"
	"time"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

type Scheduler interface {
	Start() error
	Stop()
}

type scheduler struct {
	redisCl *asynq.RedisClientOpt
	ins     *asynq.Scheduler
	cfg     *config.Config
}

func NewScheduler(cfg *config.Config, redisCl *asynq.RedisClientOpt) (Scheduler, error) {
	logger := NewLogger()

	loc, err := time.LoadLocation("Local")
	if err != nil {
		return nil, err
	}

	ins := asynq.NewScheduler(redisCl, &asynq.SchedulerOpts{
		Location: loc,
		Logger:   logger,
	})

	return &scheduler{
		redisCl: redisCl,
		ins:     ins,
		cfg:     cfg,
	}, nil
}

func (s *scheduler) Start() error {
	if err := s.initWork(); err != nil {
		return err
	}

	return s.ins.Start()
}

func (s *scheduler) Stop() {
	s.ins.Shutdown()
}

func (s *scheduler) initWork() error {
	if s.cfg.CronSpec.UpdateStatusBooking != "" {
		taskUpdateStatusBooking := asynq.NewTask(constant.TaskUpdateStatusBooking, nil, nil)
		if _, err := s.ins.Register(s.cfg.CronSpec.UpdateStatusBooking, taskUpdateStatusBooking, asynq.MaxRetry(0)); err != nil {
			return err
		}
	}
	return nil
}

func (distributor *redisTaskDistributor) DistributeTaskUpdateStatusBooking(ctx context.Context, opts ...asynq.Option) error {
	task := asynq.NewTask(constant.TaskUpdateStatusBooking, nil, nil)
	info, err := distributor.client.Enqueue(task, opts...)
	if err != nil {
		log.Error().Err(err).Msg("error when enqueue task")
		return err
	}
	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task")

	return nil
}

func (processor *redisTaskProcessor) ProcessTaskUpdateStatusBooking(ctx context.Context, task *asynq.Task) error {
	log.Info().Msg("process task update status booking")
	// get all booking that create 1 day ago
	conditions := []common.Condition{}
	conditions = append(conditions, common.Condition{
		Field:    "created_at",
		Operator: common.OperatorLessThanOrEqual,
		Value:    time.Now().AddDate(0, 0, -1),
	})
	conditions = append(conditions, common.Condition{
		Field:    "created_at",
		Operator: common.OperatorGreaterOrEqual,
		Value:    time.Now().AddDate(0, 0, -2),
	})
	conditions = append(conditions, common.Condition{
		Field:    "status_id",
		Operator: common.OperatorEqual,
		Value:    constant.BookingStatusPending,
	})

	bookings, err := processor.bookingSto.ListAllBookingWithCondition(ctx, conditions)
	if err != nil {
		log.Error().Err(err).Msg("error when get all booking with condition")
		return err
	}

	// update status booking to cancel
	for _, booking := range bookings {
		if err := processor.bookingUC.UpdateStatusBooking(ctx, booking.Id, constant.BookingStatusCancel); err != nil {
			log.Error().Err(err).Msg("error when update status booking")
			return err
		}
	}
	log.Info().Msg("update status booking success")
	return nil
}

//1    create      2           3
