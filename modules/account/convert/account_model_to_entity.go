package convert

import (
	"paradise-booking/entities"
	"paradise-booking/modules/account/iomodel"
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
		Phone:    account.Phone,
		Dob:      account.Dob,
		Avatar:   account.Avt,
		Bio:      account.Bio,
	}
}
