package entities

import "paradise-booking/common"

type BookingDetail struct {
	common.SQLModel
	BookingId       int     `json:"booking_id" gorm:"column:booking_id"`
	FullName        string  `json:"full_name" gorm:"column:full_name"`
	Phone           string  `json:"phone" gorm:"column:phone"`
	Email           string  `json:"email" gorm:"column:email"`
	Type            int     `json:"type" gorm:"column:type"`
	GuestName       string  `json:"guest_name" gorm:"column:guest_name"`
	ContentToVendor string  `json:"content_to_vendor" gorm:"column:content_to_vendor"`
	TotalPrice      float64 `json:"total_price" gorm:"column:total_price"`
	TimeTo          string  `json:"time_to" gorm:"column:time_to"`
	TimeFrom        string  `json:"time_from" gorm:"column:time_from"`
	NumberOfGuest   int     `json:"number_of_guest" gorm:"column:number_of_guest"`
	PaymentMethod   int     `json:"payment_method" gorm:"column:payment_method"`
}

func (BookingDetail) TableName() string {
	return "bookings_detail"
}
