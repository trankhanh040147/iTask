package convert

import (
	"paradise-booking/entities"
	"paradise-booking/modules/booking/iomodel"
	"paradise-booking/utils"
)

func ConvertBookingModelToBookingEntity(model *iomodel.CreateBookingReq) *entities.Booking {
	checkInTime, _ := utils.ParseStringToTime(model.CheckInDate)
	checkOutTime, _ := utils.ParseStringToTime(model.CheckOutDate)
	return &entities.Booking{
		UserId:      model.UserID,
		PlaceId:     model.PlaceID,
		CheckInDate: checkInTime,
		ChekoutDate: checkOutTime,
	}
}
