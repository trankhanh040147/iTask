package cmdworker

import (
	"github.com/hibiken/asynq"
	"iTask/config"
	accountusecase "iTask/modules/account/usecase"
	verifyemailsusecase "iTask/modules/verify_emails/usecase"
	"iTask/provider/mail"
	"iTask/worker"
	"log"
)

//func RunTaskProcessor(redisOpt *asynq.RedisClientOpt, accountSto accountusecase.AccountStorage, cfg *config.Config, verifyEmailsUC verifyemailsusecase.VerifyEmailsUseCase, bookingSto bookingusecase.BookingStorage, bookingUC bookinghandler.BookingUseCase) {
//	mailer := mail.NewGmailSender(cfg.Email.EmailSenderName, cfg.Email.EmailSenderAddress, cfg.Email.EmailSenderPassword)
//	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, accountSto, verifyEmailsUC, mailer, bookingSto, bookingUC)
//	err := taskProcessor.Start()
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func RunTaskProcessor(redisOpt *asynq.RedisClientOpt, accountSto accountusecase.AccountStorage, cfg *config.Config, verifyEmailsUC verifyemailsusecase.VerifyEmailsUseCase) {
	mailer := mail.NewGmailSender(cfg.Email.EmailSenderName, cfg.Email.EmailSenderAddress, cfg.Email.EmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, accountSto, verifyEmailsUC, mailer, nil, nil)
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal(err)
	}
}
