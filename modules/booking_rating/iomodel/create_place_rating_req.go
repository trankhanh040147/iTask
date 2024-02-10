package iomodel

type CreateBookingRatingReq struct {
	BookingID int     `json:"booking_id"`
	PlaceID   int     `json:"place_id"`
	Title     string  `json:"title"`
	Content   string  `json:"content"`
	Rating    float64 `json:"rating"`
}
