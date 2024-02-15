package iomodel

import (
	"iTask/entities"
)

type GetBookingResp struct {
	UserId  int              `json:"user_id"`
	User    entities.Account `json:"user"`
	GetData DataListBooking  `json:"data"`
}
