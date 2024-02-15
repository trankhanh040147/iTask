package iomodel

import "iTask/entities"

type ListBookingPlaceReservationResp struct {
	Data []BookingPlaceResp `json:"data"`
}

type BookingPlaceResp struct {
	*entities.Place
	IsBooked bool `json:"is_booked"`
}
