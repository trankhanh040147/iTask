package convert

import (
	"iTask/entities"
	"iTask/modules/booking/iomodel"
)

func ConvertBookingModelToBookingDetail(model *iomodel.CreateBookingReq) *entities.BookingDetail {
	return &entities.BookingDetail{
		FullName:        model.BookingInfo.FullName,
		Phone:           model.BookingInfo.Phone,
		Email:           model.BookingInfo.Email,
		Type:            model.BookingInfo.Type,
		GuestName:       model.BookingInfo.GuestName,
		ContentToVendor: model.BookingInfo.ContentToVendor,
		TotalPrice:      model.BookingInfo.TotalPrice,
		NumberOfGuest:   model.BookingInfo.NumberOfGuest,
	}
}
