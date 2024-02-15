package iomodel

import "paradise-booking/common"

type ListPlaceResp struct {
	Data   []GetPlaceResp `json:"data"`
	Paging *common.Paging `json:"paging"`
}
