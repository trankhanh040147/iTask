package paymentstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *paymentStorage) GetPaymentByVendor(ctx context.Context, vendorID int, paging *common.Paging) ([]entities.Payment, error) {
	var payments []entities.Payment

	db := s.db.Table(entities.Payment{}.TableName())

	if err := db.Raw("call GetPaymentsForVendor(?,?,?)", vendorID, paging.Page, paging.Limit).Scan(&payments).Error; err != nil {
		return nil, err
	}

	count := int64(0)
	if err := db.Raw("call GetPaymentsSizeOfVendor(?)", vendorID).Scan(&count).Error; err != nil {
		return nil, err
	}

	paging.Total = count

	return payments, nil
}
