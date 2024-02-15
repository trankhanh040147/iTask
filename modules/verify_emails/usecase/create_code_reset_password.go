package verifyemailsusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/constant"
	"paradise-booking/entities"
	"paradise-booking/utils"

	"gorm.io/gorm"
)

func (uc *verifyEmailsUseCase) UpsertResetSetCodePassword(ctx context.Context, email string) (*entities.VerifyEmail, error) {
	// create reset code password
	randomCode := utils.GenerateRandomCode(constant.LengthRandomCode)

	// set expired time
	expiredTime := utils.GetExpiredTime(constant.ExpiredTimeVerifyEmail)

	// create reset code password
	record := &entities.VerifyEmail{
		Email:     email,
		Type:      constant.TypeResetPassword,
		ScretCode: randomCode,
		ExpiredAt: &expiredTime,
	}

	// get verify email by email
	_, err := uc.verifyEmailsStore.CheckIsExist(ctx, email, constant.TypeResetPassword)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, common.ErrCannotGetEntity("verify_emails", err)
	}
	// if not exist => create
	if err == gorm.ErrRecordNotFound {
		_, err = uc.verifyEmailsStore.Create(ctx, record)
		if err != nil {
			return nil, common.ErrCannotCreateEntity("verify_emails", err)
		}
	}

	// if exist => update
	verifyEmail := entities.VerifyEmail{
		ScretCode: randomCode,
		ExpiredAt: &expiredTime,
	}
	err = uc.verifyEmailsStore.Update(ctx, email, constant.TypeResetPassword, &verifyEmail)
	if err != nil {
		return nil, common.ErrCannotUpdateEntity("verify_emails", err)
	}

	if err != nil {
		return nil, common.ErrCannotCreateEntity("verify_emails", err)
	}
	return record, nil
}
