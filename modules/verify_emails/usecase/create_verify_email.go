package verifyemailsusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/constant"
	"paradise-booking/entities"
	"paradise-booking/utils"
)

func (uc *verifyEmailsUseCase) CreateVerifyEmails(ctx context.Context, email string) (*entities.VerifyEmail, error) {
	// create verify code
	randomCode := utils.GenerateRandomCode(constant.LengthRandomCode)

	// set expired time
	expiredTime := utils.GetExpiredTime(constant.ExpiredTimeVerifyEmail)

	// create verify email
	record := &entities.VerifyEmail{
		Email:     email,
		Type:      constant.TypeVerifyEmail,
		ScretCode: randomCode,
		ExpiredAt: &expiredTime,
	}
	data, err := uc.verifyEmailsStore.Create(ctx, record)
	if err != nil {
		return nil, common.ErrCannotCreateEntity("verify_emails", err)
	}
	return data, nil
}
