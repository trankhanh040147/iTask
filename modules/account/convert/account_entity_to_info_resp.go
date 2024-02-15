package convert

import (
	"paradise-booking/entities"
	"paradise-booking/modules/account/iomodel"
)

func ConvertAccountEntityToInfoResp(account *entities.Account) *iomodel.AccountInfoResp {
	created := account.CreatedAt.Format("2006-01-02 15:04:05")
	updated := account.UpdatedAt.Format("2006-01-02 15:04:05")
	return &iomodel.AccountInfoResp{
		Id:       account.Id,
		Role:     account.Role,
		Email:    account.Email,
		Username: account.Username,
		FullName: account.FullName,
		Address:  account.Address,
		Phone:    account.Phone,
		Dob:      account.Dob,
		Avt:      account.Avatar,
		Created:  created,
		Updated:  updated,
		Bio:      account.Bio,
	}
}
