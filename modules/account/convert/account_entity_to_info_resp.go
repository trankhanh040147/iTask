package convert

import (
	"fmt"
	"iTask/entities"
	"iTask/modules/account/iomodel"
)

func ConvertAccountEntityToInfoResp(account *entities.Account) *iomodel.AccountInfoResp {
	created := account.CreatedAt.Format("2006-01-02 15:04:05")
	updated := account.UpdatedAt.Format("2006-01-02 15:04:05")

	// print log data of account
	fmt.Println("account: ", account)

	return &iomodel.AccountInfoResp{
		Id:       account.Id,
		Role:     account.Role.String(),
		Email:    account.Email,
		Username: account.Username,
		FullName: account.FullName,
		Address:  account.Address,
		Title:    account.Title,
		Phone:    account.Phone,
		Dob:      account.Dob,
		Avt:      account.Avatar,
		Cover:    account.Cover,
		Created:  created,
		Updated:  updated,
		Bio:      account.Bio,
	}
}
