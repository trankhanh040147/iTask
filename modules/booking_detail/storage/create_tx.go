package bookingdetailstorage

import (
	"context"
	"paradise-booking/entities"
)

type CreateBookingDetailTxParam struct {
	Data        *entities.BookingDetail
	AfterCreate func(data *entities.BookingDetail) error
}

func (s *bookingDetailStorage) CreateTx(ctx context.Context, createBookingDetailTxParam CreateBookingDetailTxParam) error {

	err := s.execTx(ctx, func(store *bookingDetailStorage) error {
		if err := store.Create(ctx, createBookingDetailTxParam.Data); err != nil {
			return err
		}

		if createBookingDetailTxParam.AfterCreate != nil {
			if err := createBookingDetailTxParam.AfterCreate(createBookingDetailTxParam.Data); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
