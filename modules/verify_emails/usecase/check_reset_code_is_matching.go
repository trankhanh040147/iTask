package verifyemailsusecase

import (
	"context"
	"iTask/common"
	"iTask/constant"

	"gorm.io/gorm"
)

func (uc *verifyEmailsUseCase) CheckResetCodePasswordIsMatching(ctx context.Context, email string, code string) error {
	// check if reset code password and email is matching
	data, err := uc.verifyEmailsStore.Get(ctx, email, code, constant.TypeResetPassword)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrResetCodePasswordIsNotMatching("reset code", nil)
		}
		return err
	}

	// check if reset code is expired
	if data.IsExpired() {
		return common.ErrExpiredResetCodePassword("reset code", err)
	}

	// if all is ok

	return nil
}
