package cmdworker

import (
	"paradise-booking/config"
	"paradise-booking/worker"

	"github.com/hibiken/asynq"
)

func RunTaskScheduler(redisOpt *asynq.RedisClientOpt, cfg *config.Config) {
	scheduler, err := worker.NewScheduler(cfg, redisOpt)
	if err != nil {
		panic(err)
	}

	err = scheduler.Start()
	if err != nil {
		panic(err)
	}
}
