package bookingratingusecase

import (
	"context"
	"iTask/entities"
	"iTask/modules/booking_rating/iomodel"
)

func (u *bookingRatingUsecase) MakeComment(ctx context.Context, userID int, data *iomodel.CreateBookingRatingReq) (*entities.BookingRating, error) {

	model := entities.BookingRating{
		UserId:    userID,
		BookingId: data.BookingID,
		Title:     data.Title,
		Content:   data.Content,
		Rating:    int(data.Rating),
		PlaceId:   data.PlaceID,
	}

	if _, err := u.BookingRatingSto.Create(ctx, &model); err != nil {
		return nil, err
	}

	// delete rating of place in cache
	place := entities.Place{}
	place.Id = int(data.PlaceID)
	key := place.CacheKeyPlaceRating()
	u.cache.Delete(ctx, key)

	return &model, nil
}
