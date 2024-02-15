package convert

import (
	"paradise-booking/entities"
	"paradise-booking/modules/booking/iomodel"
)

func ConvertBookingModelToResp(user *entities.Account, listDataBooking []iomodel.DataListBooking) *iomodel.ListBookingResp {
	return &iomodel.ListBookingResp{
		UserId:   user.Id,
		User:     *user,
		ListData: listDataBooking,
	}
}
