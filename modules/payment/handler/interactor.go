package paymenthandler

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

type PaymentUseCase interface {
	ListPaymentByVendorID(ctx context.Context, paging *common.Paging, vendorID int, bookingID int) ([]entities.Payment, error)
}

type paymentHandler struct {
	paymentUC PaymentUseCase
}

func NewPaymentHandler(paymentUseCase PaymentUseCase) *paymentHandler {
	return &paymentHandler{paymentUC: paymentUseCase}
}
