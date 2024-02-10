package cmdworker

import (
	"log"
	"paradise-booking/config"
	accountusecase "paradise-booking/modules/account/usecase"
	bookinghandler "paradise-booking/modules/booking/handler"
	bookingusecase "paradise-booking/modules/booking/usecase"
	verifyemailsusecase "paradise-booking/modules/verify_emails/usecase"
	"paradise-booking/provider/mail"
	"paradise-booking/worker"

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
