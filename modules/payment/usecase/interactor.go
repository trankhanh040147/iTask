package paymentusecase

import (
	"context"
	"iTask/common"
	"iTask/entities"
)

type PaymentSto interface {
	CreatePayment(ctx context.Context, payment *entities.Payment) error
	ListByCondition(ctx context.Context, condition map[string]interface{}) ([]entities.Payment, error)
	GetPaymentByVendor(ctx context.Context, vendorID int, paging *common.Paging) ([]entities.Payment, error)
}

type paymentUseCase struct {
	paymentSto PaymentSto
}

func NewPaymentUseCase(paymentSto PaymentSto) *paymentUseCase {
	return &paymentUseCase{paymentSto: paymentSto}
}
