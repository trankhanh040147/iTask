package bookinghandler

import (
	"context"
	"iTask/common"
	"iTask/entities"
	"iTask/modules/booking/iomodel"
)

type BookingUseCase interface {
	ListBooking(ctx context.Context, paging *common.Paging, filter *iomodel.FilterListBooking, userID int) (*iomodel.ListBookingResp, error)
	CreateBooking(ctx context.Context, bookingData *iomodel.CreateBookingReq) (*iomodel.CreateBookingResp, error)
	UpdateStatusBooking(ctx context.Context, bookingID, status int) error
	GetBookingByID(ctx context.Context, id int) (*iomodel.GetBookingResp, error)
	GetBookingByPlaceID(ctx context.Context, placeId int, paging *common.Paging) ([]iomodel.GetBookingByPlaceResp, error)
	ListPlaceReservationByVendor(ctx context.Context, vendorId, placeId int) (*iomodel.ListBookingPlaceReservationResp, error)
	ListBookingByCondition(ctx context.Context) ([]entities.Booking, error)
	DeleteBookingByID(ctx context.Context, id int) error
	CancelBooking(ctx context.Context, bookingID int) error
}

type bookingHandler struct {
	bookingUC BookingUseCase
}

func NewBookingHandler(bookingUC BookingUseCase) *bookingHandler {
	return &bookingHandler{bookingUC: bookingUC}
}
