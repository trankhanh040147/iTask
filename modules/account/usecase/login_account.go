package accountusecase

import (
	"context"
	"errors"
	"iTask/common"
	"iTask/entities"
	"iTask/modules/account/iomodel"
	jwtprovider "iTask/provider/jwt"
	"iTask/utils"
)

func (uc *accountUseCase) LoginAccount(ctx context.Context, accountModel *iomodel.AccountLogin) (token *jwtprovider.Token, err error) {
	// find account by email
	account, err := uc.accountStorage.GetAccountByEmail(ctx, accountModel.Email)
	if err != nil {
		return nil, common.ErrEmailNotExist(account.TableName(), err)
	}

	// check status account
	if account.Status != entities.StatusActive {
		return nil, common.ErrAccountIsNotActive(account.TableName(), errors.New("account is not active"))
	}

	// check verify account
	if account.IsEmailVerified == 0 {
		return nil, common.ErrAccountIsNotVerify(account.TableName(), errors.New("account is not verify"))
	}

	// Compare password of user with hashed password in db
	if err := utils.Compare(account.Password, accountModel.Password); err != nil {
		return nil, common.ErrEmailOrPasswordInvalid(account.TableName(), err)
	}

	// generate toke
	token, err = jwtprovider.GenerateJWT(jwtprovider.TokenPayload{Role: int(account.Role), Email: account.Email}, uc.cfg)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return token, nil
}
