package convert

import (
	"iTask/entities"
	"iTask/modules/booking/iomodel"
	"iTask/utils"
)

func ConvertBookingModelToGetResp(user *entities.Account, dataBooking *entities.Booking, place *entities.Place, bookingDetail *entities.BookingDetail) *iomodel.GetBookingResp {
	checkInTime := utils.ParseTimeToString(dataBooking.CheckInDate)
	checkOutTime := utils.ParseTimeToString(dataBooking.ChekoutDate)
	return &iomodel.GetBookingResp{
		UserId: user.Id,
		User:   *user,
		GetData: iomodel.DataListBooking{
			Id:              dataBooking.Id,
			CreatedAt:       dataBooking.CreatedAt,
			UpdatedAt:       dataBooking.UpdatedAt,
			PlaceId:         dataBooking.PlaceId,
			Place:           *place,
			StatusId:        dataBooking.StatusId,
			CheckInDate:     checkInTime,
			ChekoutDate:     checkOutTime,
			GuestName:       bookingDetail.GuestName,
			TotalPrice:      bookingDetail.TotalPrice,
			ContentToVendor: bookingDetail.ContentToVendor,
			NumberOfGuest:   bookingDetail.NumberOfGuest,
			PaymentMethod:   bookingDetail.PaymentMethod,
		},
	}
}
