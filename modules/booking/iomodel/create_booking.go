package iomodel

import "paradise-booking/entities"

type CreateBookingReq struct {
	UserID  int `json:"user_id"`
	PlaceID int `json:"place_id"`
	//StatusID     int           `json:"status_id"` // default when create booking is 1: Pending
	CheckInDate   string        `json:"checkin_date"`
	CheckOutDate  string        `json:"checkout_date"`
	PaymentMethod int           `json:"payment_method"`
	BookingInfo   BookingDetail `json:"booking_info"`
}

type BookingDetail struct {
	FullName        string  `json:"full_name"`
	Phone           string  `json:"phone"`
	Email           string  `json:"email"`
	Type            int     `json:"type"`
	GuestName       string  `json:"guest_name"`
	ContentToVendor string  `json:"content_to_vendor"`
	TotalPrice      float64 `json:"total_price"`
	NumberOfGuest   int     `json:"number_of_guest"`
}

type CreateBookingResp struct {
	BookingData entities.Booking
	PaymentUrl  string `json:"payment_url"`
}
