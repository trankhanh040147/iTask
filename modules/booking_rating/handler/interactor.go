package bookingratinghandler

import (
	"context"
	"paradise-booking/entities"
	"paradise-booking/modules/booking_rating/iomodel"
)

type placeRatingUseCase interface {
	MakeComment(ctx context.Context, userID int, data *iomodel.CreateBookingRatingReq) (*entities.BookingRating, error)
	GetCommentByPlaceID(ctx context.Context, placeID int) (*iomodel.GetCommentByPlaceResp, error)
	GetCommentByBookingID(ctx context.Context, bookingID int) ([]iomodel.GetCommentResp, error)
	GetCommentByUserID(ctx context.Context, usrID int) (*iomodel.GetCommentByUserResp, error)
	GetCommentByVendorID(ctx context.Context, vendorID int) (*iomodel.GetCommentByVendorResp, error)
	GetStatisticByPlaceID(ctx context.Context, placeId int) ([]entities.StatisticResp, error)
}

type bookingratinghandler struct {
	placeRatingUC placeRatingUseCase
}

func Newbookingratinghandler(placeRatingUseCase placeRatingUseCase) *bookingratinghandler {
	return &bookingratinghandler{placeRatingUC: placeRatingUseCase}
}
