package convert

import (
	"iTask/entities"
	"iTask/modules/account/iomodel"
)

func ConvertAccountRegisModelToEntity(account *iomodel.AccountRegister) entities.Account {
	return entities.Account{
		Email:    account.Email,
		Password: account.Password,
	}
}

func ConvertAccountLoginModelToEntity(account *iomodel.AccountLogin) entities.Account {
	return entities.Account{
		Email:    account.Email,
		Password: account.Password,
	}
}

func ConvertAccountUpdatePersonalInfoModelToEntity(account *iomodel.AccountUpdatePersonalInfo) entities.Account {
	return entities.Account{
		Username: account.Username,
		FullName: account.FullName,
		Address:  account.Address,
		Title:    account.Title,
		Phone:    account.Phone,
		Dob:      account.Dob,
		Avatar:   account.Avt,
		Cover:    account.Cover,
		Bio:      account.Bio,
	}
}
