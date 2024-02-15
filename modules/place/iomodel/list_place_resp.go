package iomodel

import "iTask/common"

type ListPlaceResp struct {
	Data   []GetPlaceResp `json:"data"`
	Paging *common.Paging `json:"paging"`
}
