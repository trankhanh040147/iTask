package iomodel

import "iTask/entities"

type GetCommentByVendorResp struct {
	ListRating []GetCommentUserByVendor
	// DataVendor *entities.Account `json:"vendor"`
}

type GetCommentUserByVendor struct {
	DataRating entities.BookingRating
	DataPlace  entities.Place   `json:"place"`
	DataUser   entities.Account `json:"user"`
}
