package convert

import (
	"paradise-booking/entities"
	"paradise-booking/modules/account/iomodel"
)

func ConvertAccountEntityToInfoMangageForAdmin(accounts []entities.Account) []iomodel.AccountInfoToAdminManageResp {
	var result []iomodel.AccountInfoToAdminManageResp
	for _, v := range accounts {
		created := v.CreatedAt.Format("2006-01-02 15:04:05")
		updated := v.UpdatedAt.Format("2006-01-02 15:04:05")

		result = append(result, iomodel.AccountInfoToAdminManageResp{
			Id:       v.Id,
			Email:    v.Email,
			Username: v.Username,
			FullName: v.FullName,
			Role:     entities.MapRole[v.Role],
			Status:   entities.MapStatus[v.Status],
			Address:  v.Address,
			Phone:    v.Phone,
			Dob:      v.Dob,
			Avt:      v.Avatar,
			Created:  created,
			Updated:  updated,
		})
	}
	return result
}
