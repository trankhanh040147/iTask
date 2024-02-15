package paymentstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *paymentStorage) ListByCondition(ctx context.Context, condition map[string]interface{}) ([]entities.Payment, error) {
	var payments []entities.Payment

	db := s.db.Table(entities.Payment{}.TableName())

	db = db.Where(condition)

	err := db.Find(&payments).Error
	if err != nil {
		return nil, err
	}

	return payments, nil
}
