package convert

import (
	"iTask/entities"
	"iTask/modules/booking/iomodel"
	"iTask/utils"
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
