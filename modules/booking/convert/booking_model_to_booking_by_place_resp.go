package convert

import (
	"paradise-booking/entities"
	"paradise-booking/modules/booking/iomodel"
	"paradise-booking/utils"
)

func ConvertBookingModelToGetByPlaceResp(user *entities.Account, dataBooking *entities.Booking, place *entities.Place, bookingDetail *entities.BookingDetail) *iomodel.GetBookingByPlaceResp {
	// parse checkin and checkout date from string to time.Time
	checkInTime := utils.ParseTimeToString(dataBooking.CheckInDate)
	checkOutTime := utils.ParseTimeToString(dataBooking.ChekoutDate)
	return &iomodel.GetBookingByPlaceResp{
		Id:              dataBooking.Id,
		CreatedAt:       dataBooking.CreatedAt,
		UpdatedAt:       dataBooking.UpdatedAt,
		UserId:          user.Id,
		User:            *user,
		PlaceId:         dataBooking.PlaceId,
		Place:           *place,
		StatusId:        dataBooking.StatusId,
		CheckInDate:     checkInTime,
		ChekoutDate:     checkOutTime,
		GuestName:       bookingDetail.GuestName,
		TotalPrice:      bookingDetail.TotalPrice,
		NumberOfGuest:   bookingDetail.NumberOfGuest,
		ContentToVendor: bookingDetail.ContentToVendor,
		PaymentMethod:   bookingDetail.PaymentMethod,
	}
}
