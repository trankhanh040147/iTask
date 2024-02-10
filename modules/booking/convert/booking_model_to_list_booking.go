package convert

import (
	"paradise-booking/entities"
	"paradise-booking/modules/booking/iomodel"
	"paradise-booking/utils"
)

func ConvertBookingModelToListBooking(data entities.Booking, place *entities.Place) iomodel.DataListBooking {
	checkInTime := utils.ParseTimeToString(data.CheckInDate)
	checkOutTime := utils.ParseTimeToString(data.ChekoutDate)
	dataBooking := iomodel.DataListBooking{
		Id:          data.Id,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		PlaceId:     data.PlaceId,
		Place:       *place,
		StatusId:    data.StatusId,
		CheckInDate: checkInTime,
		ChekoutDate: checkOutTime,
	}
	return dataBooking
}
