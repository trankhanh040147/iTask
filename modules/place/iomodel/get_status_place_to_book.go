package iomodel

type GetStatusPlaceToBookResp struct {
	NumPlaceOriginal    int                       `json:"num_place_original" form:"num_place_original"`
	NumPlaceBooked      int                       `json:"num_place_booked" form:"num_place_booked"`
	NumPlaceRemain      int                       `json:"num_place_remain" form:"num_place_remain"`
	BookingPlaceHistory []BookingPlaceHistoryResp `json:"booking_place_history" form:"booking_place_history"`
}

type BookingPlaceHistoryResp struct {
	DateFrom  string `json:"date_from" form:"date_from"`
	DateTo    string `json:"date_to" form:"date_to"`
	BookingID int    `json:"booking_id" form:"booking_id"`
}
