package iomodel

import "iTask/entities"

type GetCommentByPlaceResp struct {
	ListRating []GetCommentRespByPlace
	DataPlace  entities.Place `json:"place"`
}

type GetCommentRespByPlace struct {
	DataRating entities.BookingRating
	DataUser   entities.Account `json:"user"`
}
