package entities

import (
	"paradise-booking/common"
	"time"
)

type Booking struct {
	common.SQLModel
	UserId      int        `json:"user_id" gorm:"column:user_id"`
	PlaceId     int        `json:"place_id" gorm:"column:place_id"`
	StatusId    int        `json:"status_id" gorm:"column:status_id"`
	ChekoutDate *time.Time `json:"checkout_date" gorm:"column:checkout_date"`
	CheckInDate *time.Time `json:"checkin_date" gorm:"column:checkin_date"`
}

func (Booking) TableName() string {
	return "bookings"
}
