package paymentstorage

import (
	"context"
	"iTask/entities"
)

func (s *paymentStorage) CreatePayment(ctx context.Context, payment *entities.Payment) error {
	db := s.db.Table(payment.TableName())

	err := db.Create(payment).Error
	if err != nil {
		return err
	}

	return nil
}
