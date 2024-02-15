package verifyemailsusecase

import (
	"context"
	"iTask/common"
	"iTask/constant"
	"iTask/entities"

	"gorm.io/gorm"
)

func (uc *verifyEmailsUseCase) CheckVerifyCodeIsMatching(ctx context.Context, email string, code string) (bool, error) {
	// check if verify code and email is matching
	data, err := uc.verifyEmailsStore.Get(ctx, email, code, constant.TypeVerifyEmail)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, common.ErrVerifyCodeIsNotMatching("verify code", nil)
		}
		return false, err
	}

	// check if verify code is expired
	if data.IsExpired() {
		return true, nil
	}

	// if all is ok => update status to verified
	account := &entities.Account{
		Email:           email,
		IsEmailVerified: 1,
	}

	err = uc.accountStore.UpdateIsVerifyEmailByEmail(ctx, account.Email)
	if err != nil {
		return false, common.ErrCannotUpdateEntity("account", err)
	}

	return false, nil
}
