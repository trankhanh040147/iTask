package iomodel

import "iTask/entities"

type GetCommentResp struct {
	DataRating entities.BookingRating
	DataUser   entities.Account `json:"user"`
	DataPlace  entities.Place   `json:"place"`
}
