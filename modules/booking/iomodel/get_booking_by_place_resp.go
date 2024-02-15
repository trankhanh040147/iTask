package iomodel

import (
	"iTask/entities"
	"time"
)

type GetBookingByPlaceResp struct {
	Id              int              `json:"id" gorm:"column:id"`
	CreatedAt       *time.Time       `json:"created_at"`
	UpdatedAt       *time.Time       `json:"updated_at"`
	UserId          int              `json:"user_id"`
	User            entities.Account `json:"user"`
	PlaceId         int              `json:"place_id"`
	Place           entities.Place   `json:"place"`
	StatusId        int              `json:"status_id"`
	ChekoutDate     string           `json:"checkout_date"`
	CheckInDate     string           `json:"checkin_date"`
	GuestName       string           `json:"guest_name"`
	TotalPrice      float64          `json:"total_price"`
	ContentToVendor string           `json:"content_to_vendor"`
	NumberOfGuest   int              `json:"number_of_guest"`
	PaymentMethod   int              `json:"payment_method"`
}
