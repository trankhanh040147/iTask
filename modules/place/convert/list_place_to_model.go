package convert

import (
	"iTask/common"
	"iTask/modules/place/iomodel"
)

func ConvertPlaceToListModel(listPlace []iomodel.GetPlaceResp, paging *common.Paging) *iomodel.ListPlaceResp {
	result := &iomodel.ListPlaceResp{
		Data: listPlace,
		Paging: &common.Paging{
			Page:  paging.Page,
			Limit: paging.Limit,
			Total: paging.Total,
		},
	}
	return result
}
