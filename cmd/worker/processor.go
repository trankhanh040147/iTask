package cmdworker

import (
	"log"
	"iTask/config"
	accountusecase "iTask/modules/account/usecase"
	bookinghandler "iTask/modules/booking/handler"
	bookingusecase "iTask/modules/booking/usecase"
	verifyemailsusecase "iTask/modules/verify_emails/usecase"
	"iTask/provider/mail"
	"iTask/worker"

	"github.com/hibiken/asynq"
)

func RunTaskProcessor(redisOpt *asynq.RedisClientOpt, accountSto accountusecase.AccountStorage, cfg *config.Config, verifyEmailsUC verifyemailsusecase.VerifyEmailsUseCase, bookingSto bookingusecase.BookingStorage, bookingUC bookinghandler.BookingUseCase) {
	mailer := mail.NewGmailSender(cfg.Email.EmailSenderName, cfg.Email.EmailSenderAddress, cfg.Email.EmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, accountSto, verifyEmailsUC, mailer, bookingSto, bookingUC)
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal(err)
	}
}
