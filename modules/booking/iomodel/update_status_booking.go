package iomodel

type UpdateStatusBookingReq struct {
	BookingID int `json:"booking_id"`
	Status    int `json:"status"`
}
